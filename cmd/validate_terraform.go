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
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/di-graph/digraph/utils"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

type TerraformConfigValidatorInput struct {
	TerraformPlan             utils.ParsedTerraformPlan `json:"terraform_plan"`
	Repository                string                    `json:"repository"`
	TriggeringActionEventName string                    `json:"event_name"`
	IssueNumber               int                       `json:"issue_number"`
	CommitSHA                 string                    `json:"commit_sha"`
	Ref                       string                    `json:"ref"`
	InvocationMode            string                    `json:"invocation_mode"`
	TerraformWorkspace        string                    `json:"terraform_workspace"`
}

const validationURL = "https://app.getdigraph.com/api/validate/terraform"

func invokeDigraphValidateAPI(parsedTFPlan utils.ParsedTerraformPlan, digraphAPIKey, mode, repository, ref, commitSHA, terraformWorkspace string, issueNumber int) (utils.ValidationResponse, error) {
	requestBody := TerraformConfigValidatorInput{
		TerraformPlan:      parsedTFPlan,
		Repository:         repository,
		Ref:                ref,
		InvocationMode:     mode,
		TerraformWorkspace: terraformWorkspace,
	}

	response := utils.ValidationResponse{}

	if mode == "ci/cd" {
		if issueNumber > 0 {
			requestBody.IssueNumber = issueNumber
			requestBody.TriggeringActionEventName = "pull_request"
		} else if len(commitSHA) > 0 {
			requestBody.CommitSHA = commitSHA
			requestBody.TriggeringActionEventName = "push"
		} else {
			return response, errors.New("invalid input- must specify pull request or commit sha")
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

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Printf("cannot unmarshal api response: %s\n", err)
		return utils.ValidationResponse{}, err
	}

	return response, nil
}

func terraformRunCommand(cmd *cobra.Command) error {

	tfPlanOutput, _ := cmd.Flags().GetString("raw-output-plan")
	jsonPathPlan, _ := cmd.Flags().GetString("json-path-plan")
	tfRawPath, _ := cmd.Flags().GetString("output-path-plan")
	tfJsonOutput, _ := cmd.Flags().GetString("json-output-plan")

	terraformAPIKey, _ := cmd.Flags().GetString("terraform-api-key")
	digraphAPIKey, _ := cmd.Flags().GetString("api-key")
	repository, _ := cmd.Flags().GetString("repository")
	ref, _ := cmd.Flags().GetString("ref")

	issueNumber, _ := cmd.Flags().GetInt("issue-number")
	commitSHA, _ := cmd.Flags().GetString("commit-sha")

	mode, _ := cmd.Flags().GetString("mode")
	terraformWorkspace, _ := cmd.Flags().GetString("terraform-workspace")

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

		jsonFilePath, err = utils.FetchRemoteTerraformPlan(tfPlanOutput, terraformAPIKey)
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

		rawByteValue, _ := ioutil.ReadAll(rawOutputFile)
		jsonFilePath, err = utils.FetchRemoteTerraformPlan(string(rawByteValue), terraformAPIKey)
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

	parsedPlan, err := utils.ParseTerraformPlanJSON(jsonFilePath)
	if err != nil {
		return fmt.Errorf("error parsing JSON %s", err.Error())
	}

	output, err := invokeDigraphValidateAPI(parsedPlan, digraphAPIKey, mode, repository, ref, commitSHA, terraformWorkspace, issueNumber)
	if err != nil {
		return fmt.Errorf("error calling API %s", err.Error())
	}

	if len(tfPlanOutput) > 0 {
		// cleanup by removing temp file that was written for terraform output case
		os.Remove(jsonFilePath)
	}
	if mode == "cli" {
		utils.PrettyPrintCLIOutput(output, utils.TERRAFORM_INFRA)
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
