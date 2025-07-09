package provider

import (
	"context"

	rudderclient "github.com/cloudducoeur/terraform-rudder-provider/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNodeSettings() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNodeSettingsCreateOrUpdate,
		ReadContext:   resourceNodeSettingsRead,
		UpdateContext: resourceNodeSettingsCreateOrUpdate,
		DeleteContext: resourceNodeSettingsDelete,

		Schema: map[string]*schema.Schema{
			"node_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"properties": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceNodeSettingsCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*RudderClient)

	nodeId := d.Get("node_id").(string)
	props := make(map[string]string)

	for k, v := range d.Get("properties").(map[string]interface{}) {
		props[k] = v.(string)
	}

	settings := rudderclient.NodeSettings{
		Properties: &props,
	}

	_, _, err := client.API.NodeApi.PutNodeSettings(ctx, nodeId).NodeSettings(settings).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nodeId)
	return resourceNodeSettingsRead(ctx, d, meta)
}

func resourceNodeSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*RudderClient)
	nodeId := d.Id()

	node, _, err := client.API.NodeApi.GetNodeSettings(ctx, nodeId).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	if node.Properties != nil {
		d.Set("properties", *node.Properties)
	}

	return nil
}

func resourceNodeSettingsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Not supported: simulate deletion by clearing properties
	client := meta.(*RudderClient)
	nodeId := d.Id()

	empty := rudderclient.NodeSettings{
		Properties: &map[string]string{},
	}

	_, _, err := client.API.NodeApi.PutNodeSettings(ctx, nodeId).NodeSettings(empty).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}
