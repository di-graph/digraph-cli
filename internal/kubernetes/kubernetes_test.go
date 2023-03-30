package kubernetes

import (
	"encoding/json"
	"testing"

	kubernetes "github.com/di-graph/digraph/internal/kubernetes/test_fixtures"
	"github.com/stretchr/testify/assert"
)

func TestFormatKubernetesPlan(t *testing.T) {
	resourceChanges, err := FormatParsedKubernetesResources(kubernetes.EXAMPLE_DEPLOYMENT)
	assert.Nil(t, err)
	assert.Equal(t, len(resourceChanges), 1)
	var expectedOutput []interface{}
	json.Unmarshal([]byte(kubernetes.PARSED_JSON), &expectedOutput)
	assert.Equal(t, expectedOutput, resourceChanges)
}

func TestParseKubernetesYAMLToJSON(t *testing.T) {
	parsedPlan, err := ParseKubernetesYAMLToJSON(kubernetes.EXAMPLE_DEPLOYMENT)
	assert.Nil(t, err)
	assert.Equal(t, len(parsedPlan.ResourceChanges), 1)
	var expectedOutput []interface{}
	json.Unmarshal([]byte(kubernetes.PARSED_JSON), &expectedOutput)
	assert.Equal(t, expectedOutput, parsedPlan.ResourceChanges)
}

func TestParseKubernetesManifestList(t *testing.T) {
	parsedPlan, err := ParseKubernetesYAMLToJSON(kubernetes.EXAMPLE_LIST)
	assert.Nil(t, err)
	assert.Equal(t, len(parsedPlan.ResourceChanges), 14)

	var expectedOutput []interface{}
	err = json.Unmarshal([]byte(kubernetes.PARSED_LIST_JSON), &expectedOutput)
	assert.Nil(t, err)
	assert.Equal(t, len(expectedOutput), len(parsedPlan.ResourceChanges))
	assert.Equal(t, expectedOutput, parsedPlan.ResourceChanges)
}
