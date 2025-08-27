resource "zabbix_host" "test_host" {
  host           = var.host_name
  inventory_mode = var.host_inventory_mode

  # Interfaces
  dynamic "interfaces" {
    for_each = var.host_interfaces
    content {
      type  = interfaces.value.type
      main  = interfaces.value.main
      useip = interfaces.value.useip
      ip    = interfaces.value.ip
      dns   = interfaces.value.dns
      port  = interfaces.value.port
    }
  }

  # Groups
  dynamic "groups" {
    for_each = var.host_groups
    content {
      groupid = groups.value
    }
  }

  # TLS settings
  tls_accept       = var.host_tls_accept
  tls_connect      = var.host_tls_connect
  tls_psk_identity = var.host_psk_identity
  tls_psk          = var.host_psk

  # Tags
  dynamic "tags" {
    for_each = var.host_tags
    content {
      tag   = tags.value.tag
      value = tags.value.value
    }
  }

  # Templates
  dynamic "templates" {
    for_each = var.host_templates
    content {
      templateid = templates.value
    }
  }
}
