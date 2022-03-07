output "cert_host1" {
  value     = module.private_ca_certs.certs.host1.cert
  sensitive = true
}

output "key_host1" {
  value     = module.private_ca_certs.certs.host1.key
  sensitive = true
}
