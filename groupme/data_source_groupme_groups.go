package groupme

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	g "github.com/nhomble/groupme.go/groupme"
	"time"
)

func dataResourceGroups() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGroupsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
		},
	}
}

// get group data
func dataSourceGroupsRead(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	client := m.(*g.Client)

	all, err := client.Groups.FindAll()

	if err == nil {
		groups := make([]string, 0)
		for _, a := range all {
			groups = append(groups, a.Id)
		}
		d.SetId(time.Now().String())
		if err := d.Set("ids", groups); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	} else {
		diags = append(diags, diag.Errorf("Failed to query for groups: %s", err)...)
	}

	return diags
}
