/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/di-graph/digraph/internal/terraform"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

type TerraformConfigValidatorInput struct {
	TerraformPlan             terraform.ParsedTerraformPlan `json:"terraform_plan"`
	Repository                string                        `json:"repository"`
	TriggeringActionEventName string                        `json:"event_name"`
	IssueNumber               int                           `json:"issue_number"`
	CommitSHA                 string                        `json:"commit_sha"`
	Ref                       string                        `json:"ref"`
	InvocationMode            string                        `json:"invocation_mode"`
	TerraformWorkspace        string                        `json:"terraform_workspace"`
	GroupBy                   string                        `json:"group_by"`
	OutputFormat              string                        `json:"output_format"`
}

const validationURL = "https://app.getdigraph.com/api/validate/terraform"

func invokeDigraphValidateAPI(parsedTFPlan terraform.ParsedTerraformPlan, digraphAPIKey, mode, repository, ref, commitSHA, terraformWorkspace, groupBy, outputFormat string, issueNumber int) (string, error) {
	requestBody := TerraformConfigValidatorInput{
		TerraformPlan:      parsedTFPlan,
		Repository:         repository,
		Ref:                ref,
		InvocationMode:     mode,
		TerraformWorkspace: terraformWorkspace,
		GroupBy:            groupBy,
		OutputFormat:       outputFormat,
	}

	var response string

	if mode == "ci/cd" {
		if issueNumber > 0 {
			requestBody.IssueNumber = issueNumber
			requestBody.TriggeringActionEventName = "pull_request"
		} else if len(commitSHA) > 0 {
			requestBody.CommitSHA = commitSHA
			requestBody.TriggeringActionEventName = "push"
		} else {
			return "", errors.New("invalid input- must specify pull request or commit sha")
		}

		if len(commitSHA) > 0 && len(ref) == 0 {
			return response, errors.New("must provide branch ref associated with commit")
		}

		if len(repository) == 0 {
			return response, errors.New("must provide repository flag")
		}
	}

	requestBytes, err := json.Marshal(requestBody)
	if err != nil {
		return response, err
	}

	req, err := http.NewRequest(http.MethodPost, validationURL, bytes.NewReader(requestBytes))
	if err != nil {
		return response, err
	}

	req.Header.Set("X-Digraph-Secret-Key", digraphAPIKey)

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return response, err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		return response, fmt.Errorf("client: error with http request: status code %d with body %s", res.StatusCode, body)
	}

	return string(body), nil
}

func terraformRunCommand(cmd *cobra.Command) error {

	tfPlanOutput, _ := cmd.Flags().GetString("raw-output-plan")  // raw output of tf plan as string
	jsonPathPlan, _ := cmd.Flags().GetString("json-path-plan")   // filepath of the saved plan in json format
	tfRawPath, _ := cmd.Flags().GetString("output-path-plan")    // filepath of the saved raw output of the tf plan
	tfJsonOutput, _ := cmd.Flags().GetString("json-output-plan") // output of tf plan in json format as string

	terraformAPIKey, _ := cmd.Flags().GetString("terraform-api-key")
	digraphAPIKey, _ := cmd.Flags().GetString("api-key")
	repository, _ := cmd.Flags().GetString("repository")
	ref, _ := cmd.Flags().GetString("ref")

	issueNumber, _ := cmd.Flags().GetInt("issue-number")
	commitSHA, _ := cmd.Flags().GetString("commit-sha")

	terraformWorkspace, _ := cmd.Flags().GetString("terraform-workspace")
	groupBy, _ := cmd.Flags().GetString("group-by")
	outputFormat, _ := cmd.Flags().GetString("output-format")

	traceId, _ := cmd.Flags().GetString("traceId")
	fmt.Printf("Trace ID is %s", traceId)

	if len(digraphAPIKey) == 0 {
		err := godotenv.Load(".env")

		if err != nil {
			return fmt.Errorf("must specify api-key as argument or set it within a .env file")
		}

		digraphAPIKey = os.Getenv("api-key")
	}

	var jsonFilePath string
	var err error
	if len(tfPlanOutput) > 0 {
		if len(terraformAPIKey) == 0 {
			err := godotenv.Load(".env")

			if err != nil {
				return fmt.Errorf("must specify terraform-api-key as argument or set it within a .env file")
			}

			terraformAPIKey = os.Getenv("terraform-api-key")
		}

		jsonFilePath, err = terraform.FetchRemoteTerraformPlan(tfPlanOutput, terraformAPIKey)
		if err != nil {
			return fmt.Errorf("error getting plan json %s", err.Error())
		}
	} else if len(jsonPathPlan) > 0 {
		jsonFilePath = jsonPathPlan
	} else if len(tfRawPath) > 0 {
		if len(terraformAPIKey) == 0 {
			err := godotenv.Load(".env")

			if err != nil {
				return fmt.Errorf("must specify terraform-api-key as argument or set it within a .env file")
			}

			terraformAPIKey = os.Getenv("terraform-api-key")
		}
		rawOutputFile, err := os.Open(tfRawPath)
		if err != nil {
			return fmt.Errorf("error %s", err.Error())
		}

		rawByteValue, _ := io.ReadAll(rawOutputFile)
		jsonFilePath, err = terraform.FetchRemoteTerraformPlan(string(rawByteValue), terraformAPIKey)
		if err != nil {
			return fmt.Errorf("error getting plan json %s", err.Error())
		}
	} else if len(tfJsonOutput) > 0 {
		jsonFilePath := "/tmp/tf_plan.json"
		tempFile, err := os.Create(jsonFilePath)
		if err != nil {
			return fmt.Errorf("error getting plan json %s", err.Error())
		}

		defer tempFile.Close()
		_, err = tempFile.Write([]byte(tfJsonOutput))
		if err != nil {
			return fmt.Errorf("error getting plan json %s", err.Error())
		}
	} else {
		return fmt.Errorf("must specify raw-output-plan or json-path-plan or output-path-plan or json-output-plan")
	}

	parsedPlan, err := terraform.ParseTerraformPlanJSON(jsonFilePath)
	if err != nil {
		return fmt.Errorf("error parsing JSON %s", err.Error())
	}

	mode := "cli"

	if len(commitSHA) > 0 || issueNumber > 0 {
		mode = "ci/cd"
	}

	output, err := invokeDigraphValidateAPI(parsedPlan, digraphAPIKey, mode, repository, ref, commitSHA, terraformWorkspace, groupBy, outputFormat, issueNumber)
	if err != nil {
		return fmt.Errorf("error calling API %s", err.Error())
	}

	if len(tfPlanOutput) > 0 {
		// cleanup by removing temp file that was written for terraform output case
		os.Remove(jsonFilePath)
	}
	if mode == "cli" {
		fmt.Print(output)
	}
	return nil
}

func validateTerraform() *cobra.Command {
	tfCmd := &cobra.Command{
		Use:       "terraform",
		Short:     "Validate terraform infra config changes",
		Long:      ``,
		ValidArgs: []string{"--", "-"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return terraformRunCommand(cmd)
		},
	}

	tfCmd.Flags().String("raw-output-plan", "", "Terminal output from terraform plan command")
	tfCmd.Flags().String("output-path-plan", "", "Filepath for terminal output from terraform plan command")
	tfCmd.Flags().String("json-path-plan", "", "Filepath to terraform plan JSON file")
	tfCmd.Flags().String("json-output-plan", "", "JSON output from terraform plan command")

	tfCmd.Flags().String("terraform-api-key", "", "Terraform API Key")
	tfCmd.Flags().String("terraform-workspace", "", "Terraform workspace for associated plan")

	return tfCmd
}
