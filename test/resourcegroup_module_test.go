package test_resource_group

import (
	"testing"
	"os/exec"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
)

func TestNetworkingDemoResourceGroup(t *testing.T) {
	t.Parallel()

	out, err  := exec.Command("/bin/sh", "./scripts/az_rg_test_networking.sh" ).Output()
	if err != nil {
		t.Error("UnitTest Failed")
	}
	output := string(out[:])
	value, parseErr := strconv.Atoi(strings.TrimSpace(output))
	if parseErr != nil {
		t.Errorf("Unable to parse string to int %s", output)
	}
	assert.Equal(t, 1, value)
}

func TestAPIMDemoResourceGroup(t *testing.T) {
	t.Parallel()

	out, err  := exec.Command("/bin/sh", "./scripts/az_rg_test_apim.sh" ).Output()
	if err != nil {
		t.Error("UnitTest Failed")
	}
	output := string(out[:])
	value, parseErr := strconv.Atoi(strings.TrimSpace(output))
	if parseErr != nil {
		t.Errorf("Unable to parse string to int %s", output)
	}
	assert.Equal(t, 1, value)
}