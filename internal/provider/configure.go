package provider

import (
	"context"

	rudderclient "github.com/cloudducoeur/terraform-rudder-provider/client"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type RudderClient struct {
	API *rudderclient.APIClient
}

func configureProvider(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	cfg := rudderclient.NewConfiguration()
	cfg.Servers = rudderclient.ServerConfigurations{
		{
			URL: d.Get("api_url").(string),
		},
	}

	cfg.AddDefaultHeader("X-API-Token", d.Get("api_token").(string))

	client := rudderclient.NewAPIClient(cfg)
	return &RudderClient{API: client}, nil
}
