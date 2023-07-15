package main

import (
	"context"
	"terraform-provider-confluentacl/internal"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
	providerserver.Serve(context.Background(), internal.NewProvider, providerserver.ServeOpts{
		Address: "registry.terraform.io/bruno-zamariola/confluentacl",
		ProtocolVersion: 6,
	})
}
