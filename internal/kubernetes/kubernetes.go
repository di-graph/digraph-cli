package kubernetes

import (
	"encoding/json"
	"fmt"

	inputsanitizer "github.com/di-graph/input-sanitizer"
	"github.com/ghodss/yaml"
)

type ParsedKubernetesPlan struct {
	ResourceChanges []interface{} `json:"resource_changes"`
}

func convertYAMLToJSON(yamlString string) (interface{}, error) {
	yamlBytes := []byte(yamlString)
	jsonBytes, err := yaml.YAMLToJSON(yamlBytes)
	if err != nil {
		return "", fmt.Errorf("error converting yaml to json %s", err.Error())
	}
	var jsonOutput interface{}
	json.Unmarshal([]byte(jsonBytes), &jsonOutput)
	return jsonOutput, nil
}

func FormatParsedKubernetesPlan(manifestYaml string) ([]interface{}, error) {
	manifestJSON, err := convertYAMLToJSON(manifestYaml)
	if err != nil {
		return nil, fmt.Errorf("error converting yaml to json %s", err.Error())
	}

	var resourceChanges []interface{} = []interface{}{
		manifestJSON,
	}

	return resourceChanges, nil
}

func ParseKubernetesYAMLToJSON(kubernetesManifest string) (ParsedKubernetesPlan, error) {

	var parsedKubernetes ParsedKubernetesPlan

	resourceChanges, err := FormatParsedKubernetesPlan(kubernetesManifest)
	if err != nil {
		return ParsedKubernetesPlan{}, fmt.Errorf("error formatting parsed kubernetes plan %s", err.Error())
	}

	changeBytes, err := json.Marshal(resourceChanges)
	if err != nil {
		return ParsedKubernetesPlan{}, fmt.Errorf("error marshaling parsed plan changes %s", err.Error())
	}
	sanitizedPlan, err := inputsanitizer.SanitizeValuesByKey(changeBytes)
	if err != nil {
		return ParsedKubernetesPlan{}, fmt.Errorf("error sanitizing values %s", err.Error())
	}

	var changes []interface{}
	err = json.Unmarshal(sanitizedPlan, &changes)
	if err != nil {
		return ParsedKubernetesPlan{}, fmt.Errorf("error unmarshaling changes %s", err.Error())
	}

	parsedKubernetes.ResourceChanges = changes

	return parsedKubernetes, nil
}
