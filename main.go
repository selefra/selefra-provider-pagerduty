package main

import (
	"github.com/selefra/selefra-provider-pagerduty/provider"
	"github.com/selefra/selefra-provider-sdk/grpc/serve"
)

func main() {
	myProvider := provider.GetProvider()
	serve.Serve(myProvider.Name, myProvider)
}
