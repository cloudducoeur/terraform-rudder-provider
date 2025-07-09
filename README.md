# terraform-rudder-provider

## Usage

```hcl
provider "rudder" {
  api_url   = "https://rudder.example.com/rudder/api/latest"
  api_token = "your-api-token"
}

resource "rudder_node_settings" "example" {
  node_id = "root"
  properties = {
    "env" = "production"
    "role" = "frontend"
  }
}
```
