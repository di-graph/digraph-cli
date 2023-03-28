package kubernetes

import (
	"encoding/json"
	"testing"

	kubernetes "github.com/di-graph/digraph/internal/kubernetes/test_fixtures"
	"github.com/stretchr/testify/assert"
)

func TestFormatKubernetesPlan(t *testing.T) {
	resourceChanges, err := FormatParsedKubernetesPlan(kubernetes.EXAMPLE_DEPLOYMENT)
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
