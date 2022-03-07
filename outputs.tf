output "certs" {
  value = { for k, v in local.certs :
    k => {
      key  = tls_private_key.cert_key[k].private_key_pem
      cert = tls_locally_signed_cert.cert_pem[k].cert_pem
    }
  }
  sensitive = true
}

output "ca_pem" {
  value = tls_self_signed_cert.ca.cert_pem
}

output "ca_key" {
  value     = tls_private_key.ca.private_key_pem
  sensitive = true
}
