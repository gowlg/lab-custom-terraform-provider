package uws

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func datasourceInstance() *schema.Resource {
	return &schema.Resource{
		Read: datasourceInstanceRead,

		Schema: map[string]*schema.Schema{},
	}
}

func datasourceInstanceRead(d *schema.ResourceData, m interface{}) error {
	return nil
}
