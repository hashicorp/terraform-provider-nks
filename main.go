package main

import (
	"github.com/terraform-providers/terraform-provider-nks/nks"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: nks.Provider})
}
