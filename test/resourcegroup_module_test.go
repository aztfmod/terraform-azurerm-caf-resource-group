package test_resource_group

import (
	"testing"
    "fmt"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"encoding/json"
)

// Unit test for the CAF Resource-Group module
func TestTerraformCAFResourceGroupModule(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: ".",
	}

	//generate a tfplan output using the RunTerraformCommand function
	tfPlanOutput := "terraform.tfplan"
	terraform.Init(t, terraformOptions)
	terraform.RunTerraformCommand(t, terraformOptions,  "plan", "-out="+tfPlanOutput)

	// leverage the Terraform 0.12 feature to confert a tfplan to json
	tfjson := terraform.RunTerraformCommand(t, terraformOptions,  "show", "-json", "-no-color", tfPlanOutput)

	// convert the json data to a Golang Map
	jsonData := []byte(tfjson)
	var result map[string]interface{}
	json.Unmarshal(jsonData, &result)
	resource_changes := result["resource_changes"].([]interface{})

	// verify that we have 3 resources
	assert.Equal(t, len(resource_changes), 3)
	var nameArray []string
	for key, current_resource := range resource_changes {
		t.Log(key)
		res := current_resource.(map[string]interface{})
		change := res["change"].(map[string]interface{})
		action := change["actions"].([]interface{})
		assert.Contains(t, action, "create")
		after := change["after"].(map[string]interface{})
		assert.Contains(t, after["location"], "southeastasia")
		nameArray = append(nameArray, fmt.Sprintf("%v", after["name"]))
	}

	var expectedArray []string = []string {"unittest-apim-demo", "unittest-insights-demo", "unittest-networking-demo"}
	assert.Equal(t, nameArray, expectedArray)
}
