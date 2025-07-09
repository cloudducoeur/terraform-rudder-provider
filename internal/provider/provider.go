package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func New() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("RUDDER_API_URL", nil),
				Description: "Base URL for Rudder API",
			},
			"api_token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("RUDDER_API_TOKEN", nil),
				Description: "Rudder API token",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"rudder_node_settings": resourceNodeSettings(),
		},
		ConfigureContextFunc: configureProvider,
	}
}
