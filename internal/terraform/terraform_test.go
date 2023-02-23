package terraform_test

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/di-graph/digraph/internal/terraform"
	"github.com/stretchr/testify/assert"
)

const EXAMPLE_TF_PLAN_OUTPUT = `Running plan in the remote backend. Output will stream here. Pressing Ctrl-C
will stop streaming the logs, but will not stop the plan running remotely.

Preparing the remote plan...

To view this run in a browser, visit:
https://app.terraform.io/app/digraph/digraph-infrastructure/runs/run-y1Fbus1JCAxSPshp

Waiting for the plan to start...`

const EXAMPLE_PLAN_URL = "https://app.terraform.io/app/digraph/digraph-infrastructure/runs/run-y1Fbus1JCAxSPshp"

func TestParseTerraformURL(t *testing.T) {

	url, err := terraform.ParseTerraformURL(EXAMPLE_TF_PLAN_OUTPUT)
	assert.Nil(t, err)
	assert.Equal(t, EXAMPLE_PLAN_URL, url)
}

func TestParseTerraformRunID(t *testing.T) {
	runId, err := terraform.ParseTerraformRunID(EXAMPLE_PLAN_URL)
	assert.Nil(t, err)
	assert.Equal(t, "run-y1Fbus1JCAxSPshp", runId)
}

func TestParseTerraformPlanJSON(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	dirPath := filepath.Dir(filename)
	jsonFilePath := filepath.Join(dirPath, "test_fixtures", "soc-2-violations.json")
	parsedPlan, err := terraform.ParseTerraformPlanJSON(jsonFilePath)
	assert.Nil(t, err)
	assert.Equal(t, len(parsedPlan.ResourceChanges), 11)
	assert.Equal(t, parsedPlan.ResourceChanges[1].Name, "public_vm")
}

func TestParseTerraformPlanWrongFormatJSON(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	dirPath := filepath.Dir(filename)
	jsonFilePath := filepath.Join(dirPath, "test_fixtures", "invalid-plan-format.json")
	parsedPlan, err := terraform.ParseTerraformPlanJSON(jsonFilePath)
	assert.ErrorContains(t, err, "got unexpected file format, expected terraform json")
	fmt.Println(err.Error())
	assert.Equal(t, terraform.ParsedTerraformPlan{}, parsedPlan)
}

func TestSanitizeTerraformPlanJSON(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	dirPath := filepath.Dir(filename)
	jsonFilePath := filepath.Join(dirPath, "test_fixtures", "soc-2-violations.json")
	parsedPlan, err := terraform.ParseTerraformPlanJSON(jsonFilePath)
	assert.Nil(t, err)
	assert.Equal(t, len(parsedPlan.ResourceChanges), 11)
	assert.Equal(t, parsedPlan.ResourceChanges[1].Name, "public_vm")
	changeMap := parsedPlan.ResourceChanges[1].Change.After.(map[string]interface{})
	assert.Equal(t, true, changeMap["associate_public_ip_address"])
	assert.Equal(t, false, changeMap["get_password_data"])
	assert.Equal(t, "TEST_UNCHANGED", changeMap["kms_key_id"])
	assert.Equal(t, "REDACTED", changeMap["password"])
	assert.Equal(t, "REDACTED", changeMap["db_key"])
}
