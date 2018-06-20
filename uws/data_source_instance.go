package uws

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
)

type Instance struct {
	ID      int    `json:"id"`
	Name    string `json:"number"`
	Memory  int    `json:"memory"`
	Type    string `json:"type"`
	Tag     string `json:"tag"`
	Private bool   `json:"private"`
}

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
	var err error

	instanceID := d.Get("instance_id")
	url := fmt.Sprintf("http://localhost:3000/instances/%v", instanceID)

	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Expected status code 200 but got %v", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	var instance Instance

	err = json.Unmarshal(body, &instance)

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
