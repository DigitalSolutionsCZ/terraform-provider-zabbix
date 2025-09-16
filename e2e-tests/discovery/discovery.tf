resource "zabbix_discovery" "test_rule" {
  name            = var.discovery_name
  iprange         = var.discovery_iprange
  delay           = var.discovery_delay
  concurrency_max = var.discovery_concurrency_max
  status          = var.discovery_status

  dchecks {
    type  = var.discovery_dcheck_type
    key_  = var.discovery_dcheck_key
    ports = var.discovery_dcheck_ports
    uniq  = var.discovery_dcheck_uniq
  }
}
