package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/squarescale/terraform-provider-squarescale/squarescale"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: squarescale.Provider})
}
