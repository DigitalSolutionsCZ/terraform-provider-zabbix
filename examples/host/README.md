<!-- BEGIN_TF_DOCS -->


## Providers

| Name | Version |
|------|---------|
| <a name="provider_zabbix"></a> [zabbix](#provider\_zabbix) | n/a |

## Resources

| Name | Type |
|------|------|
| [zabbix_host.test_host](https://registry.terraform.io/providers/DigitalSolutionsCZ/zabbix/latest/docs/resources/host) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_host_groups"></a> [host\_groups](#input\_host\_groups) | n/a | `list(string)` | <pre>[<br/>  "2",<br/>  "6"<br/>]</pre> | no |
| <a name="input_host_interfaces"></a> [host\_interfaces](#input\_host\_interfaces) | n/a | <pre>list(object({<br/>    type  = number<br/>    main  = number<br/>    useip = number<br/>    ip    = string<br/>    dns   = string<br/>    port  = string<br/>  }))</pre> | <pre>[<br/>  {<br/>    "dns": "localhost",<br/>    "ip": "127.0.0.1",<br/>    "main": 1,<br/>    "port": "10050",<br/>    "type": 1,<br/>    "useip": 1<br/>  }<br/>]</pre> | no |
| <a name="input_host_inventory_mode"></a> [host\_inventory\_mode](#input\_host\_inventory\_mode) | n/a | `number` | `1` | no |
| <a name="input_host_name"></a> [host\_name](#input\_host\_name) | n/a | `string` | `"Terraform Test Host"` | no |
| <a name="input_host_psk"></a> [host\_psk](#input\_host\_psk) | n/a | `string` | `"1f87b595725ac58dd977beef14b97461a7c1045b9a1c963065002c5473194952"` | no |
| <a name="input_host_psk_identity"></a> [host\_psk\_identity](#input\_host\_psk\_identity) | n/a | `string` | `"PSK-IDENTITY-001"` | no |
| <a name="input_host_tags"></a> [host\_tags](#input\_host\_tags) | n/a | <pre>list(object({<br/>    tag   = string<br/>    value = string<br/>  }))</pre> | <pre>[<br/>  {<br/>    "tag": "Environment",<br/>    "value": "dev"<br/>  },<br/>  {<br/>    "tag": "Hosting",<br/>    "value": "internal"<br/>  }<br/>]</pre> | no |
| <a name="input_host_templates"></a> [host\_templates](#input\_host\_templates) | n/a | `list(string)` | <pre>[<br/>  "10001",<br/>  "10318"<br/>]</pre> | no |
| <a name="input_host_tls_accept"></a> [host\_tls\_accept](#input\_host\_tls\_accept) | n/a | `number` | `2` | no |
| <a name="input_host_tls_connect"></a> [host\_tls\_connect](#input\_host\_tls\_connect) | n/a | `number` | `2` | no |
| <a name="input_zabbix_api_password"></a> [zabbix\_api\_password](#input\_zabbix\_api\_password) | Zabbix password of user | `string` | `"zabbix"` | no |
| <a name="input_zabbix_api_user"></a> [zabbix\_api\_user](#input\_zabbix\_api\_user) | Zabbix user | `string` | `"Admin"` | no |
| <a name="input_zabbix_endpoint"></a> [zabbix\_endpoint](#input\_zabbix\_endpoint) | Zabbix API endpoint URL | `string` | `"http://localhost:8080/api_jsonrpc.php"` | no |
<!-- END_TF_DOCS -->