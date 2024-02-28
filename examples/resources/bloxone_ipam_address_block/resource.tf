terraform {
  required_providers {
    bloxone = {
      source  = "infobloxopen/bloxone"
      version = "0.1.9"
    }
  }
}
provider "bloxone" {
  csp_url = "https://stage.csp.infoblox.com"
  api_key = "990607e8b99bd14f7b4e4e45693b6b4eb808f2332e04afbc9bf202804829339c"
  default_tags = {
    managed_by = "terraform"
  }
}

resource "bloxone_ipam_ip_space" "example" {
  name = "example_ip_space_venkat_1"
  tags = {
    location = "site1"
  }
}

resource "bloxone_ipam_address_block" "example" {
  address = "192.168.1.0"
  cidr    = 24
  name    = "example_address_block_venkat_1"
  space   = bloxone_ipam_ip_space.example.id
}

resource "bloxone_ipam_address_block" "example_tags" {
  address = "10.0.0.0"
  cidr    = 8
  space   = bloxone_ipam_ip_space.example.id

  # Other optional fields
  name    = "example_address_block_tags_venkat_1"
  comment = "Example address block with tags created by the terraform provider"
  tags = {
    location = "site1"
  }
  asm_config = {
    asm_threshold       = 90
    enable              = "true"
    enable_notification = "true"
    forecast_period     = 10
    growth_factor       = 10
    growth_type         = "percent"
    history             = 30
    min_total           = 2
    min_unused          = 10
    reenable_date       = "2024-01-24T10:10:00+00:00"
  }
  dhcp_config = {
    allow_unknown = true
    ignore_list = [
      {
        type  = "hardware"
        value = "aa:bb:cc:dd:ee:ff"
      },
      {
        type  = "client_text"
        value = "001d.a18b.36d0"
      },
      {
        type  = "client_hex"
        value = "333964392D4769302F31"
      }
    ]
  }
}

# Next available address block
resource "bloxone_ipam_address_block" "example_na_ab" {
  next_available_id = bloxone_ipam_address_block.example.id
  cidr              = 26
  space             = bloxone_ipam_ip_space.example.id

  # Other optional fields
  name    = "example_address_block_tags_venkat_1"
  comment = "Example address block with tags created by the terraform provider"
  tags = {
    location = "site1"
  }
  asm_config = {
    asm_threshold       = 90
    enable              = "true"
    enable_notification = "true"
    forecast_period     = 10
    growth_factor       = 10
    growth_type         = "percent"
    history             = 30
    min_total           = 2
    min_unused          = 10
    reenable_date       = "2024-01-24T10:10:00+00:00"
  }
  dhcp_config = {
    allow_unknown = true
    ignore_list = [
      {
        type  = "hardware"
        value = "aa:bb:cc:dd:ee:ff"
      },
      {
        type  = "client_text"
        value = "001d.a18b.36d0"
      },
      {
        type  = "client_hex"
        value = "333964392D4769302F31"
      }
    ]
  }
}
