package uws

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceInstanceCreate,
		Read:   resourceInstanceRead,
		Update: resourceInstanceUpdate,
		Delete: resourceInstanceDelete,

		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"memory": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tag": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"private": {
				Type:     schema.TypeBool,
				Required: true,
			},
		},
	}
}

func resourceInstanceCreate(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)

	createInstanceInput := CreateInstanceInput{
		Name:    d.Get("name").(string),
		Memory:  d.Get("memory").(int),
		Type:    d.Get("type").(string),
		Tag:     d.Get("tag").(string),
		Private: d.Get("private").(bool),
	}

	uwsClient, err := NewClient(config.BaseUrl)

	if err != nil {
		return err
	}

	instance, err := uwsClient.CreateInstance(createInstanceInput)

	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%d", instance.ID))
	d.Set("instance_id", instance.ID)
	d.Set("name", instance.Name)
	d.Set("memory", instance.Memory)
	d.Set("type", instance.Type)
	d.Set("tag", instance.Tag)
	d.Set("private", instance.Private)

	return nil
}

func resourceInstanceRead(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	uwsClient, err := NewClient(config.BaseUrl)

	instance, err := uwsClient.ReadInstance(d.Get("instance_id").(int))

	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%d", instance.ID))
	d.Set("instance_id", instance.ID)
	d.Set("name", instance.Name)
	d.Set("memory", instance.Memory)
	d.Set("type", instance.Type)
	d.Set("tag", instance.Tag)
	d.Set("private", instance.Private)

	return nil
}

func resourceInstanceUpdate(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)

	instanceID := d.Get("instance_id").(int)

	updateInstanceInput := UpdateInstanceInput{
		Name:    d.Get("name").(string),
		Memory:  d.Get("memory").(int),
		Type:    d.Get("type").(string),
		Tag:     d.Get("tag").(string),
		Private: d.Get("private").(bool),
	}

	uwsClient, err := NewClient(config.BaseUrl)

	if err != nil {
		return err
	}

	instance, err := uwsClient.UpdateInstance(instanceID, updateInstanceInput)

	d.SetId(fmt.Sprintf("%d", instance.ID))
	d.Set("instance_id", instance.ID)
	d.Set("name", instance.Name)
	d.Set("memory", instance.Memory)
	d.Set("type", instance.Type)
	d.Set("tag", instance.Tag)
	d.Set("private", instance.Private)

	return nil
}

func resourceInstanceDelete(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)

	uwsClient, err := NewClient(config.BaseUrl)

	if err != nil {
		return err
	}

	err = uwsClient.DeleteInstance(d.Get("instance_id").(int))

	if err != nil {
		return err
	}

	return nil
}
