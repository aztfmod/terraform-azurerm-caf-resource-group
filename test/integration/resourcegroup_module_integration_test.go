package integrationtest_resource_group

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
	"github.com/Azure/azure-sdk-for-go/services/resourcegraph/mgmt/2019-04-01/resourcegraph"
	"context"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func TestNetworkingDemoResourceGroup(t *testing.T) {
	t.Parallel()
	//var query string = "ResourceContainers | where type =~ 'microsoft.resources/subscriptions/resourcegroups' | where name == 'test-apim-demo'| where tags['module'] =~'terraform-azurerm-caf-resource-group'" 
	var query string = "ResourceContainers | where type =~ 'microsoft.resources/subscriptions/resourcegroups' | where name == 'test-apim-demo'| where tags['module'] =~'terraform-azurerm-caf-resource-group'" 
	subscriptionID := os.Getenv("ARM_SUBSCRIPTION_ID")

	graphClient := resourcegraph.New()
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err == nil {
		graphClient.Authorizer = authorizer
	}

	queryRequest := new(resourcegraph.QueryRequest)
	queryRequest.Query = &query
	queryRequest.Subscriptions = &([]string{subscriptionID})
	result, err := graphClient.Resources(context.Background(), *queryRequest)
	if err != nil {
		t.Errorf("%+v\n", err)
	}
	assert.Equal(t, int64(1), *result.Count)
}

func TestAPIMDemoResourceGroup(t *testing.T) {
	t.Parallel()
	var query string = "ResourceContainers | where type =~ 'microsoft.resources/subscriptions/resourcegroups' | where name == 'test-apim-demo'| where tags['module'] =~'terraform-azurerm-caf-resource-group'" 
	subscriptionID := os.Getenv("ARM_SUBSCRIPTION_ID")

	graphClient := resourcegraph.New()
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err == nil {
		graphClient.Authorizer = authorizer
	}

	queryRequest := new(resourcegraph.QueryRequest)
	queryRequest.Query = &query
	queryRequest.Subscriptions = &([]string{subscriptionID})
	result, err := graphClient.Resources(context.Background(), *queryRequest)
	if err != nil {
		t.Errorf("%+v\n", err)
	}
	assert.Equal(t, int64(1), *result.Count)
}
