package main

import (
	"github.com/gowlg/lab-custom-terraform-provider/uws"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(
		&plugin.ServeOpts{
			ProviderFunc: uws.Provider,
		},
	)
}
