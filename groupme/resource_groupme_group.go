package groupme

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	g "github.com/nhomble/groupme.go/groupme"
)

func resourceGroup() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the group",
				Required:    true,
			},
			"group_id": {
				Type:        schema.TypeString,
				Description: "GroupMe internal id",
				Computed:    true,
			},
		},
		CreateContext: resourceGroupCreate,
		ReadContext:   resourceGroupRead,
		UpdateContext: resourceGroupUpdate,
		DeleteContext: resourceGroupDelete,
		Description:   "manage a groupme group",
	}
}

func resourceGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
	client := meta.(*g.Client)

	name := d.Get("name").(string)
	group, err := client.Groups.Create(&g.CreateGroupCommand{
		Name:  name,
		Share: false,
	})
	if group != nil && err == nil {
		d.Set("group_id", group.Id)
		d.SetId(group.Id)
		return diags
	} else {
		return append(diags, diag.Errorf("Failed to create group=%s error: %s", name, err)...)
	}
}

func resourceGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
	client := meta.(*g.Client)

	id, exists := d.GetOk("group_id")
	name := d.Get("name").(string)

	if exists {
		_, err := client.Groups.Update(id.(string), &g.UpdateGroupCommand{
			Name:       &name,
			Share:      false,
			OfficeMode: false,
		})

		if err != nil {
			return append(diags, diag.Errorf("Failed to read group=%s error: %s", id.(string), err)...)
		}
	} else {
		return resourceGroupCreate(ctx, d, meta)
	}

	return diags
}

func resourceGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
	client := meta.(*g.Client)

	id := d.Get("group_id").(string)

	group, err := client.Groups.Get(id)

	if group != nil && err == nil {
		d.Set("name", group.Name)
		d.SetId(group.Id)
		return diags
	} else {
		return append(diags, diag.Errorf("Failed to read group=%s error: %s", id, err)...)
	}
}

func resourceGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
	client := meta.(*g.Client)

	id := d.Get("group_id").(string)
	err := client.Groups.Delete(id)
	if err == nil {
		return diags
	} else {
		return append(diags, diag.Errorf("Failed to delete group_id=%s error: %s", id, err)...)
	}
}
