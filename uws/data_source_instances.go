package uws

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func datasourceInstances() *schema.Resource {
	return &schema.Resource{
		Read: datasourceInstancesRead,

		Schema: map[string]*schema.Schema{},
	}
}

func datasourceInstancesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}
