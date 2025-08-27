terraform {
  required_providers {
    zabbix = {
      source = "DigitalSolutionsCZ/zabbix"
    }
  }
}

provider "zabbix" {
  endpoint  = var.zabbix_endpoint
  api_token = var.zabbix_api_token
}
