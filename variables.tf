variable "organization_name" {
  description = "Org name for TLS certificates"
  default     = "Diehlabs, Inc"
}

variable "ca_common_name" {
  description = "Common name (CN) for CA certificate"
  default     = "diehlabs.com"
}

variable "csrs" {
  description = "List of attributes for CSRs to use when creating server certificates"
  type = list(object({
    dns_names    = list(string)
    ip_addresses = list(string)
    common_name  = string
  }))
}
