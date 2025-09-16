package internal

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDiscovery() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDiscoveryCreate,
		ReadContext:   resourceDiscoveryRead,
		UpdateContext: resourceDiscoveryUpdate,
		DeleteContext: resourceDiscoveryDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the discovery rule.",
			},
			"iprange": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "IP range to scan, e.g. 192.168.1.1-255.",
			},
			"delay": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Delay between checks, in seconds (as string). Default is '3600'.",
				Default:     "3600",
			},
			"status": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Rule status. 0=enabled, 1=disabled.",
				Default:     0,
			},
			"concurrency_max": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Max concurrent checks. Default is 1.",
				Default:     1,
			},
			"proxyid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the proxy used for discovery.",
			},
			"error": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Error text if there have been any problems executing the discovery rule (read-only).",
			},
			"dchecks": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "List of discovery checks for this rule.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Check type, e.g. 9=Zabbix agent.",
						},
						"key_": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Item key for the check.",
						},
						"ports": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Port(s) to use.",
						},
						"uniq": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Uniqueness flag (0 or 1).",
							Default:     0,
						},
					},
				},
			},
		},
	}
}

func resourceDiscoveryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*APIClient)

	params := map[string]interface{}{
		"name":            d.Get("name").(string),
		"iprange":         d.Get("iprange").(string),
		"delay":           d.Get("delay").(string),
		"status":          d.Get("status").(int),
		"concurrency_max": d.Get("concurrency_max").(int),
		"dchecks":         d.Get("dchecks"),
	}

	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "drule.create",
		"params":  params,
		"id":      1,
	}

	withAuth := false
	if client.AuthToken == "" {
		withAuth = true
	}

	resp, err := client.doRequest(payload, withAuth)
	if err != nil {
		return diag.Errorf("Error creating discovery rule: %s", err)
	}

	var result struct {
		Result struct {
			DruleIDs []string `json:"druleids"`
		} `json:"result"`
		Error any `json:"error"`
	}

	if err := json.Unmarshal(resp, &result); err != nil {
		return diag.Errorf("Error parsing response: %s", err)
	}

	if len(result.Result.DruleIDs) == 0 {
		return diag.Errorf("No discovery rule ID returned, error: %v", result.Error)
	}

	d.SetId(result.Result.DruleIDs[0])
	return resourceDiscoveryRead(ctx, d, m)
}

func resourceDiscoveryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*APIClient)

	druleID := d.Id()

	params := map[string]interface{}{
		"output":        "extend",
		"selectDChecks": "extend",
		"druleids":      []string{druleID},
	}

	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "drule.get",
		"params":  params,
		"id":      1,
	}

	withAuth := false
	if client.AuthToken == "" {
		withAuth = true
	}

	resp, err := client.doRequest(payload, withAuth)
	if err != nil {
		return diag.Errorf("Error reading discovery rule: %s", err)
	}

	var result struct {
		Result []map[string]interface{} `json:"result"`
		Error  any                      `json:"error"`
	}

	if err := json.Unmarshal(resp, &result); err != nil {
		return diag.Errorf("Error parsing read response: %s", err)
	}

	if len(result.Result) == 0 {
		d.SetId("")
		return nil
	}

	rule := result.Result[0]

	if v, ok := rule["name"].(string); ok {
		d.Set("name", v)
	}
	if v, ok := rule["iprange"].(string); ok {
		d.Set("iprange", v)
	}
	if v, ok := rule["delay"].(string); ok {
		d.Set("delay", v)
	}
	if v, ok := rule["proxyid"].(string); ok {
		d.Set("proxyid", v)
	}
	if v, ok := rule["error"].(string); ok {
		d.Set("error", v)
	}
	if v, ok := rule["status"].(string); ok {
		if v == "1" {
			d.Set("status", 1)
		} else {
			d.Set("status", 0)
		}
	}
	if v, ok := rule["concurrency_max"].(string); ok {
		d.Set("concurrency_max", v)
	}

	if dc, ok := rule["dchecks"].([]interface{}); ok {
		normalized := make([]map[string]interface{}, 0, len(dc))
		for _, raw := range dc {
			if check, ok := raw.(map[string]interface{}); ok {
				normalized = append(normalized, map[string]interface{}{
					"type":  check["type"],
					"key_":  check["key_"],
					"ports": check["ports"],
					"uniq":  check["uniq"],
				})
			}
		}
		d.Set("dchecks", normalized)
	}

	return nil
}

func resourceDiscoveryUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*APIClient)

	params := map[string]interface{}{
		"druleid": d.Id(),
	}

	updatableFields := []string{
		"name",
		"iprange",
		"delay",
		"status",
		"concurrency_max",
		"dchecks",
	}

	for _, field := range updatableFields {
		if d.HasChange(field) {
			params[field] = d.Get(field)
		}
	}

	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "drule.update",
		"params":  params,
		"id":      1,
	}

	withAuth := false
	if client.AuthToken == "" {
		withAuth = true
	}

	resp, err := client.doRequest(payload, withAuth)
	if err != nil {
		return diag.Errorf("Error updating discovery rule: %s", err)
	}

	var result struct {
		Result struct {
			DruleIDs []string `json:"druleids"`
		} `json:"result"`
		Error any `json:"error"`
	}

	if err := json.Unmarshal(resp, &result); err != nil {
		return diag.Errorf("Error parsing update response: %s", err)
	}

	if len(result.Result.DruleIDs) == 0 {
		return diag.Errorf("No discovery rule ID returned after update, error: %v", result.Error)
	}

	return resourceDiscoveryRead(ctx, d, m)
}

func resourceDiscoveryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*APIClient)

	druleID := d.Id()

	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "drule.delete",
		"params":  []string{druleID},
		"id":      1,
	}

	withAuth := false
	resp, err := client.doRequest(payload, withAuth)
	if err != nil {
		return diag.Errorf("Error deleting discovery rule: %s", err)
	}

	var result struct {
		Result struct {
			DruleIDs []string `json:"druleids"`
		} `json:"result"`
		Error any `json:"error"`
	}

	if err := json.Unmarshal(resp, &result); err != nil {
		return diag.Errorf("Error parsing delete response: %s", err)
	}

	if len(result.Result.DruleIDs) == 0 {
		return diag.Errorf("No discovery rule ID returned after delete, error: %v", result.Error)
	}

	d.SetId("")
	return nil
}
