package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type ValidationResponse struct {
	Comment              string                            `json:"comment"`
	ViolationsByCategory map[string]map[string]interface{} `json:"violations_by_category"`
	ViolationsByResource map[string]map[string][]string    `json:"violations_by_resource"`
}

const TERRAFORM_INFRA = "Terraform"
const KUBERNETES_INFRA = "Kubernetes"

var FRIENDLY_NAMES map[string]string = map[string]string{
	"disallowed_providers":     "Disallowed Provider(s)",
	"missing_tags":             "Missing tag(s)",
	"exposing_public_ips":      "[SOC2] Exposing Public IPs",
	"sensitive_resource_types": "Sensitive Resource Type",
}

func constructResultsByTerraformResource(output ValidationResponse) {
	numViolations := len(output.ViolationsByResource)
	if numViolations == 0 {
		color.Green(output.Comment)
	} else {
		var result string
		result = result + "\n"
		boldRed := color.New(color.FgRed, color.Bold)
		result = boldRed.Sprintf("%s%d Terraform entities have policy violations:\n", result, numViolations)

		// iterate over resources with violations and print details
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
			result = result + resourceDetailString
		}

		fmt.Println(result)
	}
}

func constructResultsByTerraformCategory(output ValidationResponse) {
	numViolations := len(output.ViolationsByResource)
	if numViolations == 0 {
		color.Green(output.Comment)
	} else {
		// iterate over categories with violations and print details
		var result string
		var numViolations int
		for category, details := range output.ViolationsByCategory {
			if len(details) > 0 {
				blueBold := color.New(color.FgBlue, color.Bold)
				friendlyName := FRIENDLY_NAMES[category]
				categoryDetailString := blueBold.Sprintf("* %s", friendlyName)
				categoryDetailString = categoryDetailString + "\n"
				for resource, detail := range details {
					underline := color.New(color.Underline)
					var joinedDetails string
					switch v := detail.(type) {
					case []string:
						if len(v) > 0 {
							joinedDetails = strings.Join(v, ", ")
						}
					case string:
						joinedDetails = v
					case bool:
						joinedDetails = strconv.FormatBool(v)
					default:
						joinedDetails = fmt.Sprintf("%v", v)
					}

					categoryDetailString = fmt.Sprintf(`
%s
    - %s: %s
			`, categoryDetailString, underline.Sprintf(resource), joinedDetails)
					numViolations++
				}
				result = result + categoryDetailString
			}
		}
		boldRed := color.New(color.FgRed, color.Bold)
		resultHeader := "\n"
		resultHeader = boldRed.Sprintf("%s%d Terraform policy violations:", resultHeader, numViolations)

		fmt.Println(resultHeader + result)
	}
}

func PrettyPrintCLIOutput(output ValidationResponse, infraType string) {
	if infraType == TERRAFORM_INFRA {
		// constructResultsByTerraformCategory(output)
		constructResultsByTerraformResource(output)
	}
}
