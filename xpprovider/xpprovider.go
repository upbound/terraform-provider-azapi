package xpprovider

import (
	"context"

	tfprovider "github.com/hashicorp/terraform-plugin-framework/provider"

	"github.com/Azure/terraform-provider-azapi/internal/provider"
)

func FrameworkProvider(_ context.Context) (tfprovider.Provider, error) {
	return provider.AzureProvider(), nil
}
