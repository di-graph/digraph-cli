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

type KubernetesManifestItem struct {
	APIVersion string      `json:"apiVersion"`
	Kind       string      `json:"kind"`
	Metadata   interface{} `json:"metadata"`
	Spec       interface{} `json:"spec"`
	Status     interface{} `json:"status"`
}

type KubernetesManifestList struct {
	APIVersion string                   `json:"apiVersion"`
	Items      []KubernetesManifestItem `json:"items"`
	Kind       string                   `json:"kind"`
	Metadata   interface{}              `json:"metadata"`
}

func convertYAMLToJSON(yamlString string) ([]byte, error) {
	yamlBytes := []byte(yamlString)
	jsonBytes, err := yaml.YAMLToJSON(yamlBytes)
	if err != nil {
		return []byte(""), fmt.Errorf("error converting yaml to json %s", err.Error())
	}
	return jsonBytes, nil
}

func parseKubernetesInput(inputJSON []byte) ([]interface{}, error) {
	var kubernetesList KubernetesManifestList
	var jsonBytes []byte
	var resourceChanges []interface{}

	err := json.Unmarshal(inputJSON, &kubernetesList)
	if err != nil {
		return resourceChanges, fmt.Errorf("error trying to unmarshal json into kubernetes manifest list %s", err.Error())
	}

	if len(kubernetesList.Items) > 0 {
		jsonBytes, err = json.Marshal(kubernetesList.Items)
		if err != nil {
			return resourceChanges, fmt.Errorf("error trying to marshal kubernetes manifest list json %s", err.Error())
		}
	} else {
		var kubernetesItem KubernetesManifestItem
		err = json.Unmarshal(inputJSON, &kubernetesItem)
		if err != nil {
			return resourceChanges, fmt.Errorf("error trying to unmarshal json into kubernetes manifest item %s", err.Error())
		}
		var changeArray []interface{} = []interface{}{
			kubernetesItem,
		}
		jsonBytes, err = json.Marshal(changeArray)
		if err != nil {
			return resourceChanges, fmt.Errorf("error trying to marshal kubernetes manifest item json %s", err.Error())
		}
	}

	err = json.Unmarshal(jsonBytes, &resourceChanges)
	if err != nil {
		return []interface{}{}, fmt.Errorf("error trying to unmarshal json into interface %s", err.Error())
	}

	return resourceChanges, nil

}

func FormatParsedKubernetesResources(manifestYaml string) ([]interface{}, error) {
	manifestJSONBytes, err := convertYAMLToJSON(manifestYaml)
	if err != nil {
		return nil, fmt.Errorf("error converting yaml to json %s", err.Error())
	}

	// check if list of deployments or single deployment
	resourceChanges, err := parseKubernetesInput(manifestJSONBytes)
	if err != nil {
		return nil, fmt.Errorf("error parsing input into valid kubernetes structs %s", err.Error())
	}

	return resourceChanges, nil
}

func ParseKubernetesYAMLToJSON(kubernetesManifest string) (ParsedKubernetesPlan, error) {

	var parsedKubernetes ParsedKubernetesPlan

	resourceChanges, err := FormatParsedKubernetesResources(kubernetesManifest)
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
