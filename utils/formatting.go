package utils

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

type ValidationResponse struct {
	Comment              string                         `json:"comment"`
	ViolationsByCategory interface{}                    `json:"violations_by_category"`
	ViolationsByResource map[string]map[string][]string `json:"violations_by_resource"`
}

const TERRAFORM_INFRA = "Terraform"
const KUBERNETES_INFRA = "Kubernetes"

var FRIENDLY_NAMES map[string]string = map[string]string{
	"disallowed_providers":     "Disallowed Provider(s)",
	"missing_tags":             "Missing tag(s)",
	"exposing_public_ips":      "[SOC2] Exposing Public IPs",
	"sensitive_resource_types": "Sensitive Resource Type",
}

func constructResultsByResource(resultString string, output ValidationResponse) string {
	// iterate over resources with violations and print details
	var result string
	for resource, details := range output.ViolationsByResource {
		blueBold := color.New(color.FgBlue, color.Bold)
		resourceDetailString := blueBold.Sprintf("* %s", resource)
		resourceDetailString = resourceDetailString + "\n"
		for category, detailList := range details {
			friendlyName := FRIENDLY_NAMES[category]
			underline := color.New(color.Underline)
			joinedDetails := strings.Join(detailList, ", ")
			resourceDetailString = fmt.Sprintf(`
%s
- %s: %s
			`, resourceDetailString, underline.Sprintf(friendlyName), joinedDetails)
		}
		result = resultString + resourceDetailString
	}
	return result
}

func PrettyPrintCLIOutput(output ValidationResponse, infraType string) {
	numViolations := len(output.ViolationsByResource)
	if numViolations == 0 {
		color.Green(output.Comment)
	} else {
		var result string
		if infraType == TERRAFORM_INFRA {
			// get total number of violating resources
			boldRed := color.New(color.FgRed, color.Bold)
			result = result + "\n"
			result = boldRed.Sprintf("%s%d Terraform entities have policy violations:\n", result, numViolations)
			result = constructResultsByResource(result, output)
		}

		fmt.Println(result)
	}
}
