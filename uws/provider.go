package uws

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"base_url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The base URL of the UWS API",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"uws_instance": datasourceInstance(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"uws_instance": resourceInstance(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		BaseUrl: d.Get("base_url").(string),
	}

	return &config, nil
}
