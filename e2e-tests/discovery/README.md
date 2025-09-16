<!-- BEGIN_TF_DOCS -->


## Providers

| Name | Version |
|------|---------|
| <a name="provider_zabbix"></a> [zabbix](#provider\_zabbix) | n/a |

## Resources

| Name | Type |
|------|------|
| [zabbix_discovery.test_rule](https://registry.terraform.io/providers/DigitalSolutionsCZ/zabbix/latest/docs/resources/discovery) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_discovery_concurrency_max"></a> [discovery\_concurrency\_max](#input\_discovery\_concurrency\_max) | Maximum number of concurrent checks (0=unlimited) | `number` | `1` | no |
| <a name="input_discovery_dcheck_key"></a> [discovery\_dcheck\_key](#input\_discovery\_dcheck\_key) | Item key for the discovery check | `string` | `"system.uname"` | no |
| <a name="input_discovery_dcheck_ports"></a> [discovery\_dcheck\_ports](#input\_discovery\_dcheck\_ports) | Port(s) to use for the discovery check | `string` | `"10050"` | no |
| <a name="input_discovery_dcheck_type"></a> [discovery\_dcheck\_type](#input\_discovery\_dcheck\_type) | Type of the discovery check (e.g. 9=Zabbix agent) | `number` | `9` | no |
| <a name="input_discovery_dcheck_uniq"></a> [discovery\_dcheck\_uniq](#input\_discovery\_dcheck\_uniq) | Uniqueness flag (0 or 1) | `number` | `0` | no |
| <a name="input_discovery_delay"></a> [discovery\_delay](#input\_discovery\_delay) | Execution interval of the discovery rule (e.g. 30s, 1m, 2h, 1d) | `string` | `"30s"` | no |
| <a name="input_discovery_iprange"></a> [discovery\_iprange](#input\_discovery\_iprange) | IP range(s) to scan, e.g. 127.0.0.1 or 192.168.1.1-255 | `string` | `"127.0.0.1"` | no |
| <a name="input_discovery_name"></a> [discovery\_name](#input\_discovery\_name) | Name of the discovery rule | `string` | `"Terraform test discovery"` | no |
| <a name="input_discovery_status"></a> [discovery\_status](#input\_discovery\_status) | Status of the rule (0=enabled, 1=disabled) | `number` | `0` | no |
| <a name="input_zabbix_api_password"></a> [zabbix\_api\_password](#input\_zabbix\_api\_password) | Zabbix password of user | `string` | `"zabbix"` | no |
| <a name="input_zabbix_api_user"></a> [zabbix\_api\_user](#input\_zabbix\_api\_user) | Zabbix user | `string` | `"Admin"` | no |
| <a name="input_zabbix_endpoint"></a> [zabbix\_endpoint](#input\_zabbix\_endpoint) | Zabbix API endpoint URL | `string` | `"http://localhost:8080/api_jsonrpc.php"` | no |
<!-- END_TF_DOCS -->