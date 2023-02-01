package cmd_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/di-graph/digraph/cmd"
	"github.com/stretchr/testify/assert"
)

func TestValidateAPIKeyArgError(t *testing.T) {
	root := cmd.RootCmd()
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs([]string{"validate", "terraform"})
	err := cmd.Execute(root)
	assert.Error(t, err)
	assert.Equal(t, "must specify api-key as argument or set it within a .env file", err.Error())
}

func TestValidateTerraformCLIMode(t *testing.T) {
	root := cmd.RootCmd()

	buf := new(bytes.Buffer)

	_, filename, _, _ := runtime.Caller(0)
	dirPath := filepath.Dir(filename)
	dirPath = filepath.Dir(dirPath)
	jsonFilePath := filepath.Join(dirPath, "internal", "terraform", "test_fixtures", "soc-2-violations.json")

	root.SetArgs([]string{"validate", "terraform", "--api-key=1234", fmt.Sprintf("--json-path-plan=%s", jsonFilePath)})

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)
		var requestPayload cmd.TerraformConfigValidatorInput
		json.Unmarshal(body, &requestPayload)
		assert.Equal(t, 11, len(requestPayload.TerraformPlan.ResourceChanges))
		assert.Equal(t, "terminal", requestPayload.OutputFormat)
		assert.Equal(t, 0, requestPayload.IssueNumber)
		assert.Equal(t, "cli", requestPayload.InvocationMode)
		w.Write([]byte("No Terraform Policy Violations"))
	}))
	defer server.Close()
	cmd.ValidationURL = server.URL
	cmd.OutputWriter = buf
	err := cmd.Execute(root)
	assert.Nil(t, err)
	out, err := ioutil.ReadAll(buf)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "No Terraform Policy Violations\n", string(out))
}

func TestValidateTerraformCICDMode(t *testing.T) {
	root := cmd.RootCmd()

	buf := new(bytes.Buffer)

	_, filename, _, _ := runtime.Caller(0)
	dirPath := filepath.Dir(filename)
	dirPath = filepath.Dir(dirPath)
	jsonFilePath := filepath.Join(dirPath, "internal", "terraform", "test_fixtures", "soc-2-violations.json")

	root.SetArgs([]string{"validate", "terraform", "--api-key=1234", fmt.Sprintf("--json-path-plan=%s", jsonFilePath), "--issue-number=1234", "--repository=testing"})

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)
		var requestPayload cmd.TerraformConfigValidatorInput
		json.Unmarshal(body, &requestPayload)
		assert.Equal(t, 11, len(requestPayload.TerraformPlan.ResourceChanges))
		assert.Equal(t, "terminal", requestPayload.OutputFormat)
		assert.Equal(t, 1234, requestPayload.IssueNumber)
		assert.Equal(t, "testing", requestPayload.Repository)
		assert.Equal(t, "ci/cd", requestPayload.InvocationMode)
		w.Write([]byte("OK"))
	}))
	defer server.Close()
	cmd.ValidationURL = server.URL
	cmd.OutputWriter = buf
	err := cmd.Execute(root)
	assert.Nil(t, err)
	out, err := ioutil.ReadAll(buf)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "", string(out))
}
