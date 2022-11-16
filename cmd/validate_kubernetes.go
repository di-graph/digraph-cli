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
	"time"

	"github.com/spf13/cobra"
)

type KubernetesConfigValidatorInput struct {
	KubernetesManifest        string `json:"kubernetes_manifest"`
	Repository                string `json:"repository"`
	TriggeringActionEventName string `json:"event_name"`
	IssueNumber               int    `json:"issue_number"`
	CommitSHA                 string `json:"commit_sha"`
	Ref                       string `json:"ref"`
	InvocationMode            string `json:"invocation_mode"`
}

const k8sValidationURL = "https://app.getdigraph.com/api/validate/kubernetes"

func invokeDigraphKubernetesValidateAPI(kubernetesManifest, digraphAPIKey, mode, repository, ref, commitSHA string, issueNumber int) (string, error) {
	requestBody := KubernetesConfigValidatorInput{
		KubernetesManifest: kubernetesManifest,
		Repository:         repository,
		Ref:                ref,
		InvocationMode:     mode,
	}

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
			return "", errors.New("must provide branch ref associated with commit")
		}

		if len(repository) == 0 {
			return "", errors.New("must provide repository flag")
		}
	}

	requestBytes, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, k8sValidationURL, bytes.NewReader(requestBytes))
	if err != nil {
		return "", err
	}

	req.Header.Set("X-Digraph-Secret-Key", digraphAPIKey)

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return "", err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		return "", errors.New(fmt.Sprintf("client: error with http request: status code %d with body %s\n", res.StatusCode, body))
	}

	return string(body), nil
}

func validateKubernetes() *cobra.Command {
	cmd := &cobra.Command{
		Use:       "kubernetes",
		Short:     "Validate kubernetes infra config changes",
		Long:      ``,
		ValidArgs: []string{"--", "-"},
		RunE: func(cmd *cobra.Command, args []string) error {
			digraphAPIKey, _ := cmd.Flags().GetString("api-key")
			repository, _ := cmd.Flags().GetString("repository")
			ref, _ := cmd.Flags().GetString("ref")

			issueNumber, _ := cmd.Flags().GetInt("issue-number")
			commitSHA, _ := cmd.Flags().GetString("commit-sha")

			mode, _ := cmd.Flags().GetString("mode")

			kubernetesManifest, _ := cmd.Flags().GetString("kubernetes-manifest")

			fmt.Println("Called kubernetes command")

			// if len(digraphAPIKey) == 0 {
			// 	err := godotenv.Load(".env")

			// 	if err != nil {
			// 		return fmt.Errorf("must specify api-key as argument or set it within a .env file")
			// 	}

			// 	digraphAPIKey = os.Getenv("api-key")
			// }

			output, err := invokeDigraphKubernetesValidateAPI(kubernetesManifest, digraphAPIKey, mode, repository, ref, commitSHA, issueNumber)
			if err != nil {
				return fmt.Errorf("error calling API %s", err.Error())
			}

			if mode == "cli" {
				fmt.Printf("%s\n", output)
			}
			return nil
		},
	}

	// cmd.Flags().String("kubernetes-manifest", "", "Kubernetes manifest YAML")

	// cmd.Flags().String("api-key", "", "Digraph API Key")

	// cmd.Flags().String("repository", "", "Github repository")

	// cmd.Flags().String("ref", "", "Branch ref")
	// cmd.Flags().Int("issue-number", 0, "Pull Request Number")
	// cmd.Flags().String("commit-sha", "", "Commit SHA")

	// cmd.Flags().String("mode", "ci/cd", "Running mode- ci/cd or cli")

	return cmd
}
