package internal

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type APIClient struct {
	Endpoint   string
	AuthToken  string
	HTTPClient http.Client
}

// Provider returns the Terraform provider schema and resources.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ZABBIX_ENDPOINT", nil),
				Description: "Full URL to Zabbix API endpoint (e.g. https://example.com/zabbix/api_jsonrpc.php).",
			},
			"api_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("ZABBIX_API_TOKEN", nil),
				Description: "Zabbix API token. Mutually exclusive with 'api_user' and 'api_password'.",
			},
			"api_user": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("ZABBIX_USER", nil),
				Description: "Zabbix API username (used with 'api_password').",
			},
			"api_password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("ZABBIX_PASSWORD", nil),
				Description: "Zabbix API password (used with 'api_user').",
			},
			"skip_ssl_verify": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ZABBIX_SKIP_SSL_VERIFY", false),
				Description: "Skip SSL certificate verification.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"zabbix_host":      resourceHost(),
			"zabbix_discovery": resourceDiscovery(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: configureProvider,
	}
}

func configureProvider(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	endpoint := d.Get("endpoint").(string)
	apiToken := d.Get("api_token").(string)
	user := d.Get("api_user").(string)
	password := d.Get("api_password").(string)
	skipSSL := d.Get("skip_ssl_verify").(bool)

	if (apiToken == "" && (user == "" || password == "")) || (apiToken != "" && (user != "" || password != "")) {
		return nil, diag.Errorf("You must specify either 'api_token' or both 'api_user' and 'api_password', but not both.")
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: skipSSL,
		},
	}
	httpClient := &http.Client{
		Transport: transport,
	}

	client := &APIClient{
		Endpoint:   endpoint,
		AuthToken:  apiToken,
		HTTPClient: *httpClient,
	}

	// If no direct API token, login with user/password to get one
	if apiToken == "" {
		loginPayload := map[string]interface{}{
			"jsonrpc": "2.0",
			"method":  "user.login",
			"params": map[string]string{
				"username": user,
				"password": password,
			},
			"id": 1,
		}

		token, err := client.callLogin(loginPayload)
		if err != nil {
			return nil, diag.FromErr(err)
		}
		client.AuthToken = token
	}

	return client, diags
}

func (c *APIClient) callLogin(payload map[string]interface{}) (string, error) {
	resp, err := c.doRequest(payload, false)
	if err != nil {
		return "", err
	}

	var result struct {
		Result string `json:"result"`
		Error  any    `json:"error"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return "", fmt.Errorf("failed to parse login response: %w", err)
	}

	if result.Result == "" {
		return "", fmt.Errorf("login failed: %v", result.Error)
	}

	return result.Result, nil
}

func (c *APIClient) doRequest(payload map[string]interface{}, withAuth bool) ([]byte, error) {
	if withAuth && c.AuthToken != "" {
		payload["auth"] = c.AuthToken
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.Endpoint, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json-rpc")
	if c.AuthToken != "" && !withAuth {
		// For direct token auth without JSON-RPC "auth" field
		req.Header.Set("Authorization", "Bearer "+c.AuthToken)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("status %d: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}
