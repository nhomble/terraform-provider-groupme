package groupme

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	g "github.com/nhomble/groupme.go/groupme"
)

// groupme provider definition
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "GroupMe API Key",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"groupme_groups": dataResourceGroups(),
			"groupme_group":  dataResourceGroup(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"groupme_group": resourceGroup(),
			"groupme_bot":   resourceBot(),
		},
		ConfigureContextFunc: configureGroupmeProvider,
	}
}

// setup groupme client
func configureGroupmeProvider(ctx context.Context, data *schema.ResourceData) (client interface{}, diags diag.Diagnostics) {
	apiKey := data.Get("api_key").(string)
	client, err := g.NewClient(g.TokenProviderFromToken(apiKey))
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	return client, diags
}
