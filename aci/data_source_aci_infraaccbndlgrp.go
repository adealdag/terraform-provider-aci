package aci

import (
	"context"
	"fmt"

	"github.com/adealdag/aci-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAciPCVPCInterfacePolicyGroup() *schema.Resource {
	return &schema.Resource{

		ReadContext: dataSourceAciPCVPCInterfacePolicyGroupRead,

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"lag_t": &schema.Schema{
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

func dataSourceAciPCVPCInterfacePolicyGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	aciClient := m.(*client.Client)

	name := d.Get("name").(string)

	rn := fmt.Sprintf("infra/funcprof/accbundle-%s", name)

	dn := fmt.Sprintf("uni/%s", rn)

	infraAccBndlGrp, err := getRemotePCVPCInterfacePolicyGroup(aciClient, dn)

	if err != nil {
		return diag.FromErr(err)
	}
	_, err = setPCVPCInterfacePolicyGroupAttributes(infraAccBndlGrp, d)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
