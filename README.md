<p align="center">
  <a href="https://registry.terraform.io/providers/DigitalSolutionsCZ/zabbix/latest/docs">
    <img src="https://camo.githubusercontent.com/cdda8928975712cecce7be8b6a1506e3b327b1643cd3391dcf40515e25b54f73/68747470733a2f2f7777772e6461746f636d732d6173736574732e636f6d2f323838352f313733313337333331302d7465727261666f726d5f77686974652e737667" alt="Terraform Logo" width="200">
  </a>
  &nbsp;&nbsp;&nbsp;
  <a href="https://github.com/DigitalSolutionsCZ/terraform-provider-zabbix">
    <img src="https://raw.githubusercontent.com/zabbix/zabbix/refs/heads/master/misc/images/docs/zabbix_logo.svg" alt="terraform-provider-zabbix" width="200">
  </a>
  &nbsp;&nbsp;&nbsp;
  <a href="https://search.opentofu.org/provider/DigitalSolutionsCZ/zabbix/latest">
    <img src="https://raw.githubusercontent.com/opentofu/brand-artifacts/main/full/transparent/SVG/on-dark.svg#gh-dark-mode-only" alt="zabbix-provider-opentofu" width="200">
  </a>
  <h3 align="center" style="font-weight: bold">Terraform Provider for Zabbix</h3>
  <p align="center">
    <a href="https://github.com/DigitalSolutionsCZ/terraform-provider-zabbix/graphs/contributors">
      <img alt="Contributors" src="https://img.shields.io/github/contributors/DigitalSolutionsCZ/terraform-provider-zabbix">
    </a>
    <a href="https://golang.org/doc/devel/release.html">
      <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/DigitalSolutionsCZ/terraform-provider-zabbix">
    </a>
    <a href="https://github.com/DigitalSolutionsCZ/terraform-provider-zabbix/actions?query=workflow%3Arelease">
      <img alt="GitHub Workflow Status" src="https://img.shields.io/github/actions/workflow/status/DigitalSolutionsCZ/terraform-provider-zabbix/release.yml?tag=latest&label=release">
    </a>
    <a href="https://github.com/DigitalSolutionsCZ/terraform-provider-zabbix/releases">
      <img alt="GitHub release (latest by date including pre-releases)" src="https://img.shields.io/github/v/release/DigitalSolutionsCZ/terraform-provider-zabbix?include_prereleases">
    </a>
  </p>
  <p align="center">
    <a href="https://github.com/DigitalSolutionsCZ/terraform-provider-zabbix/tree/main/docs"><strong>Explore the docs »</strong></a>
  </p>
</p>

# Zabbix Terraform Provider
A [Terraform](https://www.terraform.io) provider to manage [Zabbix](https://www.zabbix.com/) resources via its API using Terraform.

It supports provisioning and configuration of Zabbix users and will be extended to support other objects such as teams, stacks, endpoints, and access control.

## Requirements
- Terraform v0.13+
- Zabbix 2.x with admin API key support enabled
- Go 1.21+ (if building from source)

## Building and Installing
```hcl
make build
```

## Provider Support
| Provider                                                                                   | Provider Support Status   |
|--------------------------------------------------------------------------------------------|---------------------------|
| [Terraform](https://registry.terraform.io/providers/DigitalSolutionsCZ/zabbix/latest)      | ✅                        |
| [OpenTofu](https://search.opentofu.org/provider/DigitalSolutionsCZ/zabbix/latest)          | ✅                        |


## Example Provider Configuration

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

## Authentication

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


## Usage
See our [examples](./docs/resources/) per resources in docs.

## 🧩 Supported Resources
| Resource                                   | Documentation                                                                                  | Example                                              | Status | Terraform Import / Create => Update | E2E Tests |
|--------------------------------------------|------------------------------------------------------------------------------------------------|------------------------------------------------------|--------|-------------------------------------|-----------|
| `zabbix_host`                              | [host.md](docs/resources/host.md)                                                              | [example](examples/host/)                            | ✅     | ✅ / ✅                             | ✅        |

#### ℹ️ Note on Create ⇒ Update Behavior

Some resources support a "Create-or-Update" mechanism, when this behavior is implemented, it means:
> During the initial terraform apply, if an entity with the given name already exists, the resource will detect it and perform an update instead of attempting to create a duplicate => this is achieved by filtering existing entities by name before creation.
- This avoids the need for manual terraform import without having to have a terraform tfstate file or cleanup of existing resources in Zabbix.
- It's especially useful during migrations, initial setup, or when applying configuration into environments with pre-existing state.

---

### 💡 Missing a resource?
Is there a Zabbix resource you'd like to see supported?

👉 [Open an issue](https://github.com/DigitalSolutionsCZ/terraform-provider-zabbix/issues/new?template=feature_request.md) and we’ll consider it for implementation — or even better, submit a [Pull Request](https://github.com/DigitalSolutionsCZ/terraform-provider-zabbix/pulls) to contribute directly!

📘 See [CONTRIBUTING.md](./.github/CONTRIBUTING.md) for guidelines.

## 💬 Community & Feedback
Have questions, suggestions or want to contribute ideas?  
Want to report issues, submit pull requests or browse the source code?  
Check out the [GitHub Repository](https://github.com/DigitalSolutionsCZ/terraform-provider-zabbix) for this provider.

## ✅ Daily End-to-End Testing
To ensure maximum reliability and functionality of this provider, **automated end-to-end tests are executed every day** via GitHub Actions.

These tests run against a real Zabbix instance (started using docker compose) and validate the majority of supported resources using real Terraform plans and applies.

> 💡 This helps catch regressions early and ensures the provider remains fully operational and compatible with the Zabbix API.

### 🔄 Workflows
The project uses GitHub Actions to automate validation and testing of the provider.

- Validate and lint documentation files (`README.md` and `docs/`)
- Initialize, test and check the Zabbix provider with **Terraform** and **OpenTofu**
- Publish the new version of the Zabbix Terraform provider to Terraform Registry
- Run daily **E2E Terraform tests** against a live Zabbix instance spun up via Docker Compose (`make up`) at **07:00 UTC**

### 🧪 Localy Testing
To test the provider locally, start the Zabbix Web UI using Docker Compose:
```sh
make up
```
Then open `http://localhost:8080` in your browser.

### 🔐 Predefined Test Credentials for Login (use also E2E tests)
Thanks to the `zabbix_data` directory included in this repository, a test user and token are preloaded when you launch the local Zabbix instance:

| **Field**    | **Value**                                                                  |
|--------------|----------------------------------------------------------------------------|
| Username     | `Admin`                                                                    |
| Password     | `zabbix`                                                                   |
| API Token    | `d4719154c1a1891852acd2e2948e3f9bcdd98319d251c746c610c491cb6bba04`         |

You can now apply your Terraform templates and observe changes live in the UI.

### Testing a new version of the Zabbix provider
After making changes to the provider source code, follow these steps:
Build the provider binary:
```sh
make build
```
Install the binary into the local Terraform plugin directory:
```sh
make install-plugin
```
Update your main.tf to use the local provider source
Add the following to your Terraform configuration:
```sh
terraform {
  required_providers {
    zabbix = {
      source  = "localdomain/local/zabbix"
    }
  }
}
```
Now you're ready to test your provider against the local Zabbix instance.

## Roadmap
See the [open issues](https://github.com/DigitalSolutionsCZ/terraform-provider-zabbix/issues) for a list of proposed features (and known issues). See [CONTRIBUTING](./.github/CONTRIBUTING.md) for more information.

## License
This module is 100% Open Source and is distributed under the MIT License.  
See the [LICENSE](https://github.com/DigitalSolutionsCZ/terraform-provider-zabbix/blob/main/LICENSE) file for more information.


## Acknowledgements
- [HashiCorp Terraform](https://www.hashicorp.com/products/terraform)
- [Zabbix](https://www.zabbix.com/)
- [OpenTofu](https://opentofu.org/)
- [Docker](https://www.docker.com/)
