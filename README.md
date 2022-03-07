# terraform-<provider>-<module name>

This repo will be used as a template for new Terraform module Github repos.

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
| <a name="output_ca_key"></a> [ca\_key](#output\_ca\_key) | n/a |
| <a name="output_ca_pem"></a> [ca\_pem](#output\_ca\_pem) | n/a |
| <a name="output_certs"></a> [certs](#output\_certs) | n/a |
<!-- END_TF_DOCS -->
