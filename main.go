package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/ilya-lesikov/terraform-provider-gitfile/gitfile"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: gitfile.Provider,
	})
}
