package unittest_resource_group

import (
	"testing"
	"os"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
	"github.com/hashicorp/terraform-json"
)

// UnitTesting Terraform Plan result. This test parses the json output of Terraform Plan
func TestTerraformPlan(t *testing.T) {
	t.Parallel()
	const JSON_ENV = "JSON_PLAN_PATH"
	// Open Json file
	jsonFilePath := os.Getenv(JSON_ENV)
	if len(jsonFilePath) == 0 {
		assert.Fail(t, "JSON_PLAN_PATH env variable is empty")
	}
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		t.Errorf("%+v\n", err)
		assert.Fail(t, "Unable to read tfplan.json")
	}

	defer jsonFile.Close()
	
	//Reading Json file 
    byteValue, _ := ioutil.ReadAll(jsonFile)

	tfPlan := tfjson.Plan{}
	planErr := tfPlan.UnmarshalJSON(byteValue)
	if  planErr != nil {
		t.Errorf("%+v\n", planErr)
		assert.Fail(t, "Unable to read unmarshal tfPlan")
	}
	var resourcesChangeCount int = len(tfPlan.ResourceChanges)
	assert.Equal(t, 3, resourcesChangeCount, "Expected resources count are the same")
	
	var apimDemoRG = tfPlan.ResourceChanges[0].Change.After.(map[string]interface{})
	assert.Equal(t, true, tfPlan.ResourceChanges[0].Change.Actions.Create(), "TF Action is set to Create a new resource for test-apim-demo")
	assert.Equal(t, "test-apim-demo", apimDemoRG["name"], "Resource name is valid for test-apim-demo")
	
	var insightDemoRG = tfPlan.ResourceChanges[1].Change.After.(map[string]interface{})
	assert.Equal(t, true, tfPlan.ResourceChanges[1].Change.Actions.Create(), "TF Action is set to Create a new resource for test-insights-demo")
	assert.Equal(t, "test-insights-demo", insightDemoRG["name"], "Resource name is valid for test-insights-demo")
	
	var networkingDemoRG = tfPlan.ResourceChanges[2].Change.After.(map[string]interface{})
	assert.Equal(t, true, tfPlan.ResourceChanges[2].Change.Actions.Create(), "TF Action is set to Create a new resource for test-networking-demo")
	assert.Equal(t, "test-networking-demo", networkingDemoRG["name"], "Resource name is valid for test-networking-demo")
}
