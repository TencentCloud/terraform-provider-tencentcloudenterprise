package main

import (
	"context"
	"flag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"log"
	"terraform-provider-tencentcloudenterprise/tencentcloud"
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debuggable", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	if debugMode {
		err := plugin.Debug(context.Background(), "registry.terraform.io/tencentcloud/tencentcloudenterprise",
			&plugin.ServeOpts{
				ProviderFunc: tencentcloud.Provider,
			})
		if err != nil {
			log.Println(err.Error())
		}
	} else {
		plugin.Serve(&plugin.ServeOpts{
			ProviderFunc: tencentcloud.Provider})
	}

	//app := &cli.App{}
	//app.Run(os.Args)
}
