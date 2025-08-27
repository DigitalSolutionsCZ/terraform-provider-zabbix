# 🌐 Terraform Resource: `zabbix_host`

The `zabbix_host` resource allows you to manage hosts in Zabbix, including their interfaces, groups, templates, macros, tags, and inventory properties.

---

## Example Usage

### Basic host

```hcl
resource "zabbix_host" "example" {
  host = "example-host"
  name = "Example Host"
  status = 0  # 0=enabled, 1=disabled

  groups = [
    { groupid = "2" }
  ]

  interfaces = [
    {
      type = 1
      main = 1
      useip = 1
      ip = "192.168.1.10"
      dns = ""
      port = "10050"
    }
  ]
}
```

### Host with templates, tags, and macros

```hcl
resource "zabbix_host" "example_full" {
  host = "example-host-full"
  name = "Example Full Host"
  description = "Host with templates, tags and macros"
  status = 0

  groups = [
    { groupid = "2" }
  ]

  templates = [
    { templateid = "10001" }
  ]

  tags = [
    { tag = "env", value = "prod" },
    { tag = "role", value = "web" }
  ]

  macros = [
    { macro = "{$MACRO1}", value = "value1" }
  ]
}
```

### Host with TLS and proxy configuration

```hcl
resource "zabbix_host" "secure_host" {
  host = "secure-host"
  name = "Secure Host"

  groups = [
    { groupid = "2" }
  ]

  interfaces = [
    { type = 1, main = 1, useip = 1, ip = "10.0.0.5", port = "10050" }
  ]

  tls_connect = 4       # 1=No encryption, 2=PSK, 4=certificate
  tls_accept = 4
  tls_psk_identity = "identity123"
  tls_psk = "secretpsk"

  monitored_by = 1
  proxyid = 3
  proxy_groupid = 5
}
```

---

## Lifecycle & Behavior

* Hosts are created, updated, and deleted via the Zabbix API.
* Updating any attribute (e.g., `name`, `interfaces`, `templates`) will trigger an update request.
* To delete a host, simply run:

```hcl
terraform destroy
```

* To update a host, change attributes and re-apply:

```hcl
terraform apply
```

---

## Arguments Reference

| Name               | Type         | Required                  | Description                                                                                                               |
| ------------------ | ------------ | ------------------------- | ------------------------------------------------------------------------------------------------------------------------- |
| `host`             | string       | ✅ yes                     | Technical host name (unique in Zabbix).                                                                                   |
| `name`             | string       | 🚫 optional               | Visible name of the host. Defaults to `host` if not set.                                                                  |
| `description`      | string       | 🚫 optional               | Host description.                                                                                                         |
| `status`           | int          | 🚫 optional (default `0`) | Host status: 0=enabled, 1=disabled.                                                                                       |
| `interfaces`       | list(object) | ✅ yes                     | List of host interfaces. Each interface requires `type` and `main`, optional `useip`, `ip`, `dns`, `port`, and `details`. |
| `groups`           | list(object) | ✅ yes                     | Host groups to assign. Each object requires `groupid`.                                                                    |
| `templates`        | list(object) | 🚫 optional               | Templates to link. Each object requires `templateid`.                                                                     |
| `tags`             | list(object) | 🚫 optional               | Host tags. Each object requires `tag` and `value`.                                                                        |
| `macros`           | list(object) | 🚫 optional               | User macros. Each object requires `macro` and `value`, optional `description`.                                            |
| `inventory`        | map          | 🚫 optional               | Host inventory properties.                                                                                                |
| `inventory_mode`   | int          | 🚫 optional (default `0`) | Inventory mode: 0=manual, 1=automatic.                                                                                    |
| `tls_connect`      | int          | 🚫 optional               | TLS connection mode: 1=No encryption, 2=PSK, 4=certificate.                                                               |
| `tls_accept`       | int          | 🚫 optional               | TLS accept mode: 1=No encryption, 2=PSK, 4=certificate.                                                                   |
| `tls_psk_identity` | string       | 🚫 optional               | Pre-shared key identity for PSK encryption.                                                                               |
| `tls_psk`          | string       | 🚫 optional               | Pre-shared key value for PSK encryption (write-only).                                                                     |
| `monitored_by`     | int          | 🚫 optional               | Source monitoring the host: 0=Zabbix server, 1=Proxy, 2=Proxy group.                                                      |
| `proxyid`          | int          | 🚫 optional               | ID of the proxy monitoring the host.                                                                                      |
| `proxy_groupid`    | int          | 🚫 optional               | ID of the proxy group monitoring the host.                                                                                |

---

## Attributes Reference

| Name | Description           |
| ---- | --------------------- |
| `id` | ID of the Zabbix host |

> Attributes are populated after creation or read.
