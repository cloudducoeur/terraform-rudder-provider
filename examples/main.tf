terraform {
  required_providers {
    rudder = {
      source = "local/cloudducoeur/rudder"
      version = "1.0.0"
    }
  }
}

provider "rudder" {
  api_url   = var.rudder_api_url
  api_token = var.rudder_api_token
}

# Exemple de ressource pour tester le provider
resource "rudder_node_settings" "example" {
  node_id = var.node_id
  
  properties = {
    "env"        = "dev"
    "datacenter" = "paris"
    "owner"      = "team-infra"
    "project"    = "rudder-terraform-test"
  }
}
