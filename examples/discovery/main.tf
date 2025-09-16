terraform {
  required_providers {
    zabbix = {
      source = "DigitalSolutionsCZ/zabbix"
    }
  }
}

provider "zabbix" {
  endpoint     = var.zabbix_endpoint
  api_user     = var.zabbix_api_user
  api_password = var.zabbix_api_password
}
