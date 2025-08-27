# Zabbix Terraform Provider
A [Terraform](https://www.terraform.io) provider to manage [Zabbix](https://www.zabbix.com/) resources via its API using Terraform.

It supports provisioning and configuration of Zabbix users and will be extended to support other objects such as hosts, templates, triggers, users etc.

## 🏷️ Provider Support
| Provider       | Provider Support Status              |
|----------------|--------------------------------------|
| [Terraform](https://registry.terraform.io/providers/DigitalSolutionsCZ/zabbix/latest)      | ![Done](https://img.shields.io/badge/status-done-brightgreen)           |
| [OpenTofu](https://search.opentofu.org/provider/DigitalSolutionsCZ/zabbix/latest)       | ![Done](https://img.shields.io/badge/status-done-brightgreen) |

## ⚙️ Example Provider Configuration

```hcl
provider "zabbix" {
  endpoint = "https://zabbix.example.com/zabbix/api_jsonrpc.php"

  # Option 1: API token authentication
  api_token  = "your-api-token"

  # Option 2: Username/password authentication (retrieves auth token via user.login internally)
  # api_user     = "Admin"
  # api_password = "your-password"

  skip_ssl_verify  = true # optional (default value is `false`)
}
```

## 🔐 Authentication

The Zabbix Terraform provider supports two authentication methods:

1. **API Token** (via `Authorization: Bearer <token>` header)
2. **Username & Password** (via `user.login` → short-lived auth token internally used)

Only one method is required – if both are provided, `api_token` takes precedence.

#### Usage – API Token:

```hcl
provider "zabbix" {
  endpoint   = "https://zabbix.example.com/zabbix/api_jsonrpc.php"
  api_token  = "your-api-token"
}
```

#### Usage – Username & Password:

```hcl
provider "zabbix" {
  endpoint     = "https://zabbix.example.com/zabbix/api_jsonrpc.php"
  api_user     = "Admin"
  api_password = "your-password"
}
```

### Environment variables

You can also configure the provider via environment variables:

#### API token method

```bash
export ZABBIX_ENDPOINT="https://zabbix.example.com/zabbix/api_jsonrpc.php"
export ZABBIX_API_TOKEN="your-api-token"
export ZABBIX_SKIP_SSL_VERIFY=true
```

#### Username and password method

```bash
export ZABBIX_ENDPOINT="https://zabbix.example.com/zabbix/api_jsonrpc.php"
export ZABBIX_USER="Admin"
export ZABBIX_PASSWORD="your-password"
export ZABBIX_SKIP_SSL_VERIFY=true
```

## Arguments Reference
| Name              | Type    | Required | Description                                                                                          |
| ----------------- | ------- | -------- | ---------------------------------------------------------------------------------------------------- |
| `endpoint`        | string  | ✅ yes    | Full URL of the Zabbix API endpoint (e.g. `https://host/zabbix/api_jsonrpc.php`).                    |
| `api_token`       | string  | ❌ no     | API token for authentication. Mutually exclusive with `api_user` and `api_password`.                 |
| `api_user`        | string  | ❌ no     | Username for authentication (must be used with `api_password`). Mutually exclusive with `api_token`. |
| `api_password`    | string  | ❌ no     | Password for authentication (must be used with `api_user`). Mutually exclusive with `api_token`.     |
| `skip_ssl_verify` | boolean | ❌ no     | Skip TLS certificate verification (useful for self-signed certs). Default: `false`.                  |

## 🧩 Supported Resources
| Resource                                       | Status                                                                |
|------------------------------------------------|-----------------------------------------------------------------------|
| `zabbix_host`                                  | ![Done](https://img.shields.io/badge/status-done-brightgreen)         |

---

### 💡 Missing a resource?
Is there a Zabbix resource you'd like to see supported?

👉 [Open an issue](https://github.com/DigitalSolutionsCZ/terraform-provider-zabbix/issues/new?template=feature_request.md) and we’ll consider it for implementation — or even better, submit a [Pull Request](https://github.com/DigitalSolutionsCZ/terraform-provider-zabbix/pulls) to contribute directly!

📘 See [CONTRIBUTING.md](https://github.com/DigitalSolutionsCZ/terraform-provider-zabbix/blob/main/.github/CONTRIBUTING.md) for guidelines.

## 💬 Community & Feedback
Have questions, suggestions or want to contribute ideas?  
Want to report issues, submit pull requests or browse the source code?  
Check out the [GitHub Repository](https://github.com/DigitalSolutionsCZ/terraform-provider-zabbix) for this provider.

## ✅ Daily End-to-End Testing
To ensure maximum reliability and functionality of this provider, **automated end-to-end tests are executed every day** via GitHub Actions.

These tests run against a real Zabbix instance (started using docker compose) and validate the majority of supported resources using real Terraform plans and applies.

> 💡 This helps catch regressions early and ensures the provider remains fully operational and compatible with the Zabbix API.

## License
This module is 100% Open Source and is distributed under the MIT License.  
See the [LICENSE](https://github.com/DigitalSolutionsCZ/terraform-provider-zabbix/blob/main/LICENSE) file for more information.


## Acknowledgements
- [HashiCorp Terraform](https://www.hashicorp.com/products/terraform)
- [Zabbix](https://www.zabbix.com/)
- [OpenTofu](https://opentofu.org/)
- [Docker](https://www.docker.com/)
