module "private_ca_certs" {
  source = "../.."
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
