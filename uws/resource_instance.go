package uws

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceInstanceCreate,
		Read:   resourceInstanceRead,
		Update: resourceInstanceUpdate,
		Delete: resourceInstanceDelete,

		Schema: map[string]*schema.Schema{},
	}
}

func resourceInstanceCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInstanceRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInstanceUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInstanceDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
