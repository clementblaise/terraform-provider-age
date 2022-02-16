package provider

import (
	"context"

	"filippo.io/age"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSecretKey() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "The resource `age_secret_key` generates an Age secret/public key pair.",

		CreateContext: resourceAgeSecretKeyCreate,
		ReadContext:   resourceAgeSecretKeyRead,
		DeleteContext: resourceAgeSecretKeyDelete,

		Schema: map[string]*schema.Schema{
			"secret_key": {
				Type:        schema.TypeString,
				Description: "Generated Bech32 secret key in ASCII format.",
				Computed:    true,
				Sensitive:   true,
			},
			"public_key": {
				Type:        schema.TypeString,
				Description: "Generated Bech32 public key in ASCII format.",
				Computed:    true,
			},
		},
	}
}

func resourceAgeSecretKeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	k, err := age.GenerateX25519Identity()
	if err != nil {
		diag.FromErr(err)
	}
	tflog.Trace(ctx, "created a resource")

	d.SetId(k.Recipient().String())

	err = d.Set("secret_key", k.String())
	if err != nil {
		diag.FromErr(err)
	}

	err = d.Set("public_key", k.Recipient().String())
	if err != nil {
		diag.FromErr(err)
	}

	return nil
}

func resourceAgeSecretKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceAgeSecretKeyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceAgeSecretKeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}
