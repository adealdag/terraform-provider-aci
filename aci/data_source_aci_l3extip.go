package aci

import (
	"context"
	"fmt"

	"github.com/adealdag/aci-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAciL3outPathAttachmentSecondaryIp() *schema.Resource {
	return &schema.Resource{

		ReadContext: dataSourceAciL3outPathAttachmentSecondaryIpRead,

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"l3out_path_attachment_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"addr": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ipv6_dad": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"name_alias": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		}),
	}
}

func dataSourceAciL3outPathAttachmentSecondaryIpRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	aciClient := m.(*client.Client)

	addr := d.Get("addr").(string)

	rn := fmt.Sprintf("addr-[%s]", addr)
	LeafPortDn := d.Get("l3out_path_attachment_dn").(string)

	dn := fmt.Sprintf("%s/%s", LeafPortDn, rn)

	l3extIp, err := getRemoteL3outPathAttachmentSecondaryIp(aciClient, dn)

	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(dn)
	_, err = setL3outPathAttachmentSecondaryIpAttributes(l3extIp, d)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
