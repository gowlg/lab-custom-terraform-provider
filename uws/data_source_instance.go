package uws

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
)

func datasourceInstance() *schema.Resource {
	return &schema.Resource{
		Read: datasourceInstanceRead,

		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"memory": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"private": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func datasourceInstanceRead(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)

	instanceID := d.Get("instance_id").(int)

	uwsClient, err := NewClient(config.BaseUrl)

	if err != nil {
		return err
	}

	instance, err := uwsClient.ReadInstance(instanceID)

	if err != nil {
		return err
	}

	d.SetId(time.Now().UTC().String())
	d.Set("instance_id", instance.ID)
	d.Set("name", instance.Name)
	d.Set("memory", instance.Memory)
	d.Set("type", instance.Type)
	d.Set("tag", instance.Tag)
	d.Set("private", instance.Private)

	return nil
}
