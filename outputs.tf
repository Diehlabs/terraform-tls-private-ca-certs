output "certs" {
  description = <<EOD
Map of certficates
hostname = {
  key = private_key_pem_string
  cert = cert_pem_string
}
EOD

  value = { for k, v in local.certs :
    k => {
      key  = tls_private_key.cert_key[k].private_key_pem
      cert = tls_locally_signed_cert.cert_pem[k].cert_pem
    }
  }

  sensitive = true
}

output "ca_pem" {
  description = "The CA certificate PEM string"
  value       = tls_self_signed_cert.ca.cert_pem
}

output "ca_key" {
  description = "The CA certificate key PEM string"
  value       = tls_private_key.ca.private_key_pem
  sensitive   = true
}
