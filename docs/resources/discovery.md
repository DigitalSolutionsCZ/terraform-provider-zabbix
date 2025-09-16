# 🌐 Terraform Resource: `zabbix_discovery`

The `zabbix_discovery` resource allows you to manage **network discovery rules** in Zabbix, including IP ranges, concurrency, proxy assignment, and discovery checks.

---

## Example Usage

### Basic discovery rule

```hcl
resource "zabbix_discovery" "example" {
  name    = "Zabbix agent discovery"
  iprange = "192.168.1.1-255"

  dchecks {
    type  = 9
    key_  = "system.uname"
    ports = "10050"
    uniq  = 0
  }
}
```

### Discovery rule with proxy and custom delay

```hcl
resource "zabbix_discovery" "proxy_rule" {
  name            = "Discover via proxy"
  iprange         = "10.0.0.1-255"
  delay           = "5m"
  concurrency_max = 5
  status          = 0
  proxyid         = "10255" # Proxy ID from Zabbix

  dchecks {
    type  = 3
    key_  = ""
    ports = "21"
  }

  dchecks {
    type  = 4
    key_  = ""
    ports = "80"
  }
}
```

---

## Lifecycle & Behavior

* Discovery rules are created, updated, and deleted via the Zabbix API (`drule.create`, `drule.update`, `drule.delete`).
* Updating any attribute (e.g., `iprange`, `dchecks`) will trigger an update request.
* To delete a discovery rule, simply run:

```hcl
terraform destroy
```

* To update a discovery rule, change attributes and re-apply:

```hcl
terraform apply
```

* Existing discovery rules can be imported by ID:

```bash
terraform import zabbix_discovery.example 6
```

---

## Arguments Reference

| Name              | Type         | Required                   | Description                                                                     |
| ----------------- | ------------ | -------------------------- | ------------------------------------------------------------------------------- |
| `name`            | string       | ✅ yes                      | Name of the discovery rule.                                                     |
| `iprange`         | string       | ✅ yes                      | IP range(s) to check, comma-separated (e.g. `192.168.1.1-255,192.168.2.1-255`). |
| `delay`           | string       | 🚫 optional (default `1h`) | Execution interval, e.g. `30s`, `1m`, `2h`, `1d`, or a user macro.              |
| `status`          | int          | 🚫 optional (default `0`)  | Status of the rule: 0=enabled, 1=disabled.                                      |
| `concurrency_max` | int          | 🚫 optional (default `0`)  | Max number of concurrent checks: 0=unlimited, 1=one, 2–999=custom number.       |
| `proxyid`         | string       | 🚫 optional                | ID of the proxy used for discovery.                                             |
| `dchecks`         | list(object) | ✅ yes                      | List of discovery checks. Each block requires at least `type` and `key_`.       |
| └─ `type`         | int          | ✅ yes                      | Type of check (e.g. 9=Zabbix agent, 3=FTP, 4=HTTP).                             |
| └─ `key_`         | string       | ✅ yes                      | Item key for the check (e.g. `system.uname`).                                   |
| └─ `ports`        | string       | 🚫 optional                | Ports to use for the check (e.g. `"10050"` or `"80,443"`).                      |
| └─ `uniq`         | int          | 🚫 optional (default `0`)  | Uniqueness flag.                                                                |

---

## Attributes Reference

| Name    | Description                                                              |
| ------- | ------------------------------------------------------------------------ |
| `id`    | ID of the discovery rule (`druleid`).                                    |
| `error` | Error text if there have been any problems executing the discovery rule. |

> Attributes are populated after creation or read.
