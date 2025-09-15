package xpprovider

import (
	"context"

	tfprovider "github.com/hashicorp/terraform-plugin-framework/provider"

	"github.com/Azure/terraform-provider-azapi/internal/provider"
	"github.com/Azure/terraform-provider-azapi/internal/services/parse"
)

func FrameworkProvider(_ context.Context) (tfprovider.Provider, error) {
	return provider.AzureProvider(), nil
}

func DataPlaneResourceId(name, parentId, resourceType string) (string, error) {
	id, err := parse.NewDataPlaneResourceId(name, parentId, resourceType)
	if err != nil {
		return "", err
	}
	return id.ID(), nil
}

func NewResourceID(name, parentId, resourceType string) (string, error) {
	id, err := parse.NewResourceID(name, parentId, resourceType)
	if err != nil {
		return "", err
	}
	return id.ID(), nil
}

func ResourceIDWithResourceType(azureResourceId, resourceType string) (string, error) {
	id, err := parse.ResourceIDWithResourceType(azureResourceId, resourceType)
	if err != nil {
		return "", err
	}
	return id.ID(), nil
}

func UpdateResourceID(name, parentId, azureResourceId, resourceType string) (string, error) {
	if azureResourceId != "" {
		return ResourceIDWithResourceType(azureResourceId, resourceType)
	}
	return NewResourceID(name, parentId, resourceType)
}
