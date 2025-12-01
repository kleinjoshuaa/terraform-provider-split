---
layout: "split"
page_title: "Provider: Split"
sidebar_current: "docs-split-index"
description: |-
  The Split provider is used to interact with resources provided by the Split API.
---

# Split Provider

The Split provider is used to interact with resources provided by the
[Split API](https://docs.split.io/reference#introduction).

## Contributing

Development happens in the [GitHub repo](https://github.com/davidji99/terraform-provider-split):

* [Releases](https://github.com/davidji99/terraform-provider-split/releases)
* [Issues](https://github.com/davidji99/terraform-provider-split/issues)

## Example Usage

```hcl
provider "split" {
}

# Create a new Split environment
resource "split_environment" "foobar" {
  # ...
}
```

## Authentication

The Split provider offers a flexible means of providing credentials for authentication.
The following methods are supported, listed in order of precedence, and explained below:

- Static credentials
- Environment variables

The provider supports three authentication methods:

1. **API Key Authentication (Default)**: Uses a Bearer token in the Authorization header.
2. **Harness Token Authentication**: Uses the `x-api-key` header for authentication.
3. **Harness Platform API Key Authentication**: Uses the `x-api-key` header for authentication (same as Harness Token).

### Static credentials

Credentials can be provided statically by adding authentication arguments to the Split provider block:

```hcl
provider "split" {
  # Option 1: API Key (Bearer token authentication)
  api_key = "SOME_API_KEY"
  
  # Option 2: Harness Token (x-api-key header authentication)
  # harness_token = "SOME_HARNESS_TOKEN"
  
  # Option 3: Harness Platform API Key (x-api-key header authentication, takes precedence over harness_token)
  # harness_platform_api_key = "SOME_HARNESS_PLATFORM_API_KEY"
}
```

### Environment variables

When the Split provider block does not contain authentication arguments, the missing credentials will be sourced
from the environment variables. The provider checks in the following order:

1. `HARNESS_PLATFORM_API_KEY` (highest precedence)
2. `HARNESS_TOKEN`
3. `SPLIT_API_KEY` (lowest precedence)

> **Note**: The `HARNESS_PLATFORM_API_KEY` environment variable is shared with the [Harness Terraform provider](https://registry.terraform.io/providers/harness/harness/latest/docs). When using both providers in the same Terraform configuration, you can set this single environment variable to authenticate both providers, simplifying your CI/CD pipeline setup.

```hcl
provider "split" {}
```

```shell
# Option 1: Using Harness Platform API Key (highest precedence)
# This works for both Split and Harness providers when used together
$ export HARNESS_PLATFORM_API_KEY="SOME_HARNESS_PLATFORM_API_KEY"
$ terraform plan

# Option 2: Using Harness Token
$ export HARNESS_TOKEN="SOME_HARNESS_TOKEN"
$ terraform plan

# Option 3: Using Split API Key (default)
$ export SPLIT_API_KEY="SOME_API_KEY"
$ terraform plan
```

### Using with Harness Provider

When using both the Split and Harness Terraform providers together, you can use the shared `HARNESS_PLATFORM_API_KEY` environment variable:

```hcl
provider "harness" {
  # Uses HARNESS_PLATFORM_API_KEY environment variable
}

provider "split" {
  # Also uses HARNESS_PLATFORM_API_KEY environment variable
}
```

```shell
# Set once, works for both providers
$ export HARNESS_PLATFORM_API_KEY="your-platform-api-key"
$ terraform plan
```

### Authentication Method Notes

When using **Harness Token** or **Harness Platform API Key** authentication (x-api-key header), the following resources are deprecated and cannot be used:

- `split_user`
- `split_group`
- `split_workspace`
- `split_api_key` (only when `type = "admin"`)

## Rate Limiting

The Split provider provides automatic backoff in the event the provider detects the Split API has
[rate limited](https://docs.split.io/reference/rate-limiting) your Terraform operations. Please note
this backoff has a max timeout that can be configured by [`client_timeout`](#argument-reference) within
your `provider {}` block.

## Argument Reference

The following arguments are supported:

* `api_key` - (Optional) Split API key for Bearer token authentication. It can be provided directly or
  sourced from the `SPLIT_API_KEY` environment variable. Required if `harness_token` or `harness_platform_api_key` are not set.

* `harness_token` - (Optional) Harness token for x-api-key header authentication. It can be provided directly or
  sourced from the `HARNESS_TOKEN` environment variable. Takes precedence over `api_key` if both are set.

* `harness_platform_api_key` - (Optional) Harness Platform API key for x-api-key header authentication. It can be provided directly or
  sourced from the `HARNESS_PLATFORM_API_KEY` environment variable. Takes precedence over both `harness_token` and `api_key` if multiple are set.
  **Note**: The `HARNESS_PLATFORM_API_KEY` environment variable is shared with the Harness Terraform provider, allowing you to use a single environment variable when working with both providers.

* `base_url` - (Optional) Custom API URL.
  Can also be sourced from the `SPLIT_API_URL` environment variable.

* `remove_environment_from_state_only` - (Optional) Configure `split_environment` to only remove the resource from
  state upon deletion. This is to address out-of-band, UI based prerequisites Split has when deleting an environment.
  Defaults to `false`.

* `client_timeout` - (Optional) Configure client (http) timeout before aborting. This is to address the client retrying forever.
  It's expressed in an integer that represents seconds. Defaults to `300` seconds, or `5` minutes.