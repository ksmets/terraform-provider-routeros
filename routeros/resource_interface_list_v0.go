package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceInterfaceListV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			MetaResourcePath: PropResourcePath("/interface/list"),
			MetaId:           PropId(Name),

			"builtin": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"comment": PropCommentRw,
			"dynamic": PropDynamicRo,
			"exclude": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"include": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": PropNameForceNewRw,
		},
	}
}
