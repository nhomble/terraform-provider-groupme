package groupme

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	g "github.com/nhomble/groupme.go/groupme"
)

func resourceBot() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the bot",
				Required:    true,
			},
			"group_id": {
				Type:        schema.TypeString,
				Description: "The id of the group assigned to the bot",
				Required:    true,
			},
			"avatar_url": {
				Type:        schema.TypeString,
				Description: "URL for the bot picture",
			},
			"callback_url": {
				Type:        schema.TypeString,
				Description: "Endpoint that GroupMe will invoke on every new message in the group",
			},
			"bot_id": {
				Type:        schema.TypeString,
				Description: "GroupMe identifier for the bot",
				Computed:    true,
			},
		},
		CreateContext: resourceBotCreate,
		ReadContext:   resourceBotRead,
		UpdateContext: resourceBotUpdate,
		DeleteContext: resourceBotDelete,
		Description:   "manage a groupme bot",
	}
}

func resourceBotCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
	client := meta.(*g.Client)
	name := d.Get("name").(string)
	group_id := d.Get("group_id").(string)
	avatar_url := d.Get("avatar_url").(string)
	callback_url := d.Get("callback_url").(string)
	bot, err := client.Bots.Create(g.CreateBotCommand{
		Name:        name,
		GroupId:     group_id,
		AvatarURL:   &avatar_url,
		CallbackURL: &callback_url,
	})

	if bot != nil && err == nil {
		d.Set("bot_id", bot.BotId)
		d.SetId(bot.BotId)
		return diags
	} else {
		return append(diags, diag.Errorf("Failed to create bot=%s error: %s", name, err)...)
	}
}

func resourceBotUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
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

func resourceBotRead(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
	client := meta.(*g.Client)
	bots, err := client.Bots.List()
	id := d.Get("bot_id").(string)

	group, err := client.Groups.Get(id)

	if group != nil && err == nil {
		d.Set("name", group.Name)
		d.SetId(group.Id)
		return diags
	} else {
		return append(diags, diag.Errorf("Failed to read group=%s error: %s", id, err)...)
	}
}

func resourceBotDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
	client := meta.(*g.Client)

	id := d.Get("group_id").(string)
	err := client.Groups.Delete(id)
	if err == nil {
		return diags
	} else {
		return append(diags, diag.Errorf("Failed to delete group_id=%s error: %s", id, err)...)
	}
}
