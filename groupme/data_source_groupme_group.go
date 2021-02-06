package groupme

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	g "github.com/nhomble/groupme.go/groupme"
)

func dataResourceGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGroupRead,
		Schema: map[string]*schema.Schema{
			"group_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "GroupMe internal id",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    false,
				Description: "Name of the group",
				Computed:    true,
			},
			"image_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Default:     nil,
				Description: "GroupMe Image Service URL for the group icon",
			},
		},
	}
}

// get group data
func dataSourceGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	client := m.(*g.Client)

	id := d.Get("group_id").(string)

	group, err := client.Groups.Get(id)

	if group != nil && err == nil {
		d.Set("name", group.Name)
		d.Set("image_url", group.ImageUrl)
		d.SetId(group.Id)
		return diags
	} else {
		return append(diags, diag.Errorf("Failed to read group=%s error: %s", id, err)...)
	}
}
