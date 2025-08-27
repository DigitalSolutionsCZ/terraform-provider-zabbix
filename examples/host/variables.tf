variable "zabbix_endpoint" {
  type        = string
  default     = "http://localhost:8080/api_jsonrpc.php"
  description = "Zabbix API endpoint URL"
}

variable "zabbix_api_user" {
  type        = string
  default     = "Admin"
  description = "Zabbix user"
  sensitive   = true
}

variable "zabbix_api_password" {
  type        = string
  default     = "zabbix"
  description = "Zabbix password of user"
  sensitive   = true
}

variable "host_name" {
  type    = string
  default = "Terraform Test Host"
}

variable "host_inventory_mode" {
  type    = number
  default = 1
}

variable "host_interfaces" {
  type = list(object({
    type  = number
    main  = number
    useip = number
    ip    = string
    dns   = string
    port  = string
  }))
  default = [
    {
      type  = 1
      main  = 1
      useip = 1
      ip    = "127.0.0.1"
      dns   = "localhost"
      port  = "10050"
    }
  ]
}

variable "host_groups" {
  type    = list(string)
  default = ["2", "6"]
}

variable "host_tls_accept" {
  type    = number
  default = 2
}

variable "host_tls_connect" {
  type    = number
  default = 2
}

variable "host_psk_identity" {
  type    = string
  default = "PSK-IDENTITY-001"
}

variable "host_psk" {
  type    = string
  default = "1f87b595725ac58dd977beef14b97461a7c1045b9a1c963065002c5473194952"
}

variable "host_tags" {
  type = list(object({
    tag   = string
    value = string
  }))
  default = [
    { tag = "Environment", value = "dev" },
    { tag = "Hosting", value = "internal" }
  ]
}

variable "host_templates" {
  type    = list(string)
  default = ["10001", "10318"]
}
