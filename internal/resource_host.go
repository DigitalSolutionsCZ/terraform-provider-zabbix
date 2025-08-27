package internal

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHost() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceHostCreate,
		ReadContext:   resourceHostRead,
		UpdateContext: resourceHostUpdate,
		DeleteContext: resourceHostDelete,
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Technical name of the host.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Visible name of the host. Defaults to `host` if not set.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the host.",
			},
			"status": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "Status of the host. 0=enabled, 1=disabled.",
			},
			"interfaces": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "List of host interfaces.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"main": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"useip": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  1,
						},
						"ip": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"dns": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"port": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "10050",
						},
						"details": {
							Type:     schema.TypeMap,
							Optional: true,
						},
					},
				},
			},
			"groups": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Host groups to assign.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"groupid": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"templates": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Templates to link.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"templateid": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"tags": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Host tags.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tag": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"macros": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "User macros.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"macro": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"inventory": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Host inventory properties.",
			},
			"inventory_mode": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			// TLS parameters
			"tls_connect": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "TLS connection mode. 1=No encryption, 2=PSK, 4=certificate.",
			},
			"tls_accept": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "TLS accept mode. 1=No encryption, 2=PSK, 4=certificate.",
			},
			"tls_psk_identity": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Pre-shared key identity for PSK encryption.",
			},
			"tls_psk": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Pre-shared key value for PSK encryption (write-only).",
				Sensitive:   true,
			},
			// Proxy / monitored_by
			"monitored_by": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Source monitoring the host. 0=Zabbix server, 1=Proxy, 2=Proxy group.",
			},
			"proxyid": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "ID of the proxy monitoring the host.",
			},
			"proxy_groupid": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "ID of the proxy group monitoring the host.",
			},
		},
	}
}

func resourceHostCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*APIClient)

	params := map[string]interface{}{
		"host":   d.Get("host").(string),
		"groups": d.Get("groups"),
	}
	optionalFields := []string{
		"name", "description", "status",
		"interfaces", "templates", "tags", "macros",
		"inventory", "inventory_mode",
		"tls_connect", "tls_accept", "tls_psk_identity", "tls_psk",
		"monitored_by", "proxyid", "proxy_groupid",
	}

	for _, field := range optionalFields {
		if v, ok := d.GetOk(field); ok {
			params[field] = v
		}
	}

	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "host.create",
		"params":  params,
		"id":      1,
	}

	withAuth := false
	if client.AuthToken == "" {
		withAuth = true
	}

	resp, err := client.doRequest(payload, withAuth)
	if err != nil {
		return diag.Errorf("Error creating host: %s", err)
	}

	var result struct {
		Result struct {
			HostIDs []string `json:"hostids"`
		} `json:"result"`
		Error any `json:"error"`
	}

	if err := json.Unmarshal(resp, &result); err != nil {
		return diag.Errorf("Error parsing response: %s", err)
	}

	if len(result.Result.HostIDs) == 0 {
		return diag.Errorf("No host ID returned, error: %v", result.Error)
	}

	d.SetId(result.Result.HostIDs[0])
	return resourceHostRead(ctx, d, m)
}

func resourceHostRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// TODO: implement read via host.get
	return nil
}

func resourceHostUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// TODO: implement update via host.update
	return resourceHostRead(ctx, d, m)
}

func resourceHostDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*APIClient)

	hostID := d.Id()

	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "host.delete",
		"params":  []string{hostID},
		"id":      1,
	}

	withAuth := false
	resp, err := client.doRequest(payload, withAuth)
	if err != nil {
		return diag.Errorf("Error deleting host: %s", err)
	}

	var result struct {
		Result struct {
			HostIDs []string `json:"hostids"`
		} `json:"result"`
		Error any `json:"error"`
	}

	if err := json.Unmarshal(resp, &result); err != nil {
		return diag.Errorf("Error parsing delete response: %s", err)
	}

	if len(result.Result.HostIDs) == 0 {
		return diag.Errorf("No host ID returned after delete, error: %v", result.Error)
	}

	d.SetId("")
	return nil
}
