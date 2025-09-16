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

variable "discovery_name" {
  description = "Name of the discovery rule"
  type        = string
  default     = "Terraform test discovery"
}

variable "discovery_iprange" {
  description = "IP range(s) to scan, e.g. 127.0.0.1 or 192.168.1.1-255"
  type        = string
  default     = "127.0.0.1"
}

variable "discovery_delay" {
  description = "Execution interval of the discovery rule (e.g. 30s, 1m, 2h, 1d)"
  type        = string
  default     = "30s"
}

variable "discovery_concurrency_max" {
  description = "Maximum number of concurrent checks (0=unlimited)"
  type        = number
  default     = 1
}

variable "discovery_status" {
  description = "Status of the rule (0=enabled, 1=disabled)"
  type        = number
  default     = 0
}

variable "discovery_dcheck_type" {
  description = "Type of the discovery check (e.g. 9=Zabbix agent)"
  type        = number
  default     = 9
}

variable "discovery_dcheck_key" {
  description = "Item key for the discovery check"
  type        = string
  default     = "system.uname"
}

variable "discovery_dcheck_ports" {
  description = "Port(s) to use for the discovery check"
  type        = string
  default     = "10050"
}

variable "discovery_dcheck_uniq" {
  description = "Uniqueness flag (0 or 1)"
  type        = number
  default     = 0
}
