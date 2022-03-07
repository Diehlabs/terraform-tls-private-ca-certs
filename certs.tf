locals {
  certs = { for csr in var.csrs :
    (csr.common_name) => {
      dns_names    = csr.dns_names
      ip_addresses = csr.ip_addresses
    }
  }
}

resource "tls_private_key" "cert_key" {
  for_each  = local.certs
  algorithm = "RSA"
  rsa_bits  = "4096"
}

resource "tls_cert_request" "cert_csr" {
  for_each        = local.certs
  key_algorithm   = tls_private_key.cert_key[each.key].algorithm
  private_key_pem = tls_private_key.cert_key[each.key].private_key_pem
  dns_names = concat(
    each.value.dns_names,
    ["localhost"]
  )
  ip_addresses = concat(
    each.value.ip_addresses,
    ["127.0.0.1"]
  )
  subject {
    common_name  = each.key
    organization = var.organization_name
  }
}

resource "tls_locally_signed_cert" "cert_pem" {
  for_each              = local.certs
  cert_request_pem      = tls_cert_request.cert_csr[each.key].cert_request_pem
  ca_key_algorithm      = tls_private_key.ca.algorithm
  ca_private_key_pem    = tls_private_key.ca.private_key_pem
  ca_cert_pem           = tls_self_signed_cert.ca.cert_pem
  validity_period_hours = 2400
  allowed_uses = [
    "key_encipherment",
    "digital_signature",
    "client_auth",
    "server_auth",
  ]
}
