# terraform-tls-private-ca-certs

Generates a local CA and one or more TLS keypairs for servers.

Example usage:
```hcl
module "private_ca_certs" {
  source = ".modules/tls"
  csrs = [
    {
      common_name  = "host1"
      dns_names    = ["alias1"]
      ip_addresses = ["192.168.13.10"]
    },
    {
      common_name = "host2"
      dns_names   = ["alias2"]
      ip_addresses = [
        "192.168.13.11",
        "192.168.13.12"
      ]
    }
  ]
}
```

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_tls"></a> [tls](#requirement\_tls) | ~> 3.1.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_tls"></a> [tls](#provider\_tls) | 3.1.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [tls_cert_request.cert_csr](https://registry.terraform.io/providers/hashicorp/tls/latest/docs/resources/cert_request) | resource |
| [tls_locally_signed_cert.cert_pem](https://registry.terraform.io/providers/hashicorp/tls/latest/docs/resources/locally_signed_cert) | resource |
| [tls_private_key.ca](https://registry.terraform.io/providers/hashicorp/tls/latest/docs/resources/private_key) | resource |
| [tls_private_key.cert_key](https://registry.terraform.io/providers/hashicorp/tls/latest/docs/resources/private_key) | resource |
| [tls_self_signed_cert.ca](https://registry.terraform.io/providers/hashicorp/tls/latest/docs/resources/self_signed_cert) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_ca_common_name"></a> [ca\_common\_name](#input\_ca\_common\_name) | Common name (CN) for CA certificate | `string` | `"diehlabs.com"` | no |
| <a name="input_csrs"></a> [csrs](#input\_csrs) | List of attributes for CSRs to use when creating server certificates | <pre>list(object({<br>    dns_names    = list(string)<br>    ip_addresses = list(string)<br>    common_name  = string<br>  }))</pre> | n/a | yes |
| <a name="input_organization_name"></a> [organization\_name](#input\_organization\_name) | Org name for TLS certificates | `string` | `"Diehlabs, Inc"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_ca_key"></a> [ca\_key](#output\_ca\_key) | The CA certificate key PEM string |
| <a name="output_ca_pem"></a> [ca\_pem](#output\_ca\_pem) | The CA certificate PEM string |
| <a name="output_certs"></a> [certs](#output\_certs) | Map of certficates<br>hostname = {<br>  key = private\_key\_pem\_string<br>  cert = cert\_pem\_string<br>} |
<!-- END_TF_DOCS -->
