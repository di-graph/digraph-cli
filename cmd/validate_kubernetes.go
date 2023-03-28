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

	"github.com/di-graph/digraph/internal/error_handling"
	"github.com/di-graph/digraph/internal/kubernetes"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

type KubernetesConfigValidatorInput struct {
	KubernetesManifest        kubernetes.ParsedKubernetesPlan `json:"kubernetes_manifest"`
	Repository                string                          `json:"repository"`
	TriggeringActionEventName string                          `json:"event_name"`
	IssueNumber               int                             `json:"issue_number"`
	CommitSHA                 string                          `json:"commit_sha"`
	Ref                       string                          `json:"ref"`
	InvocationMode            string                          `json:"invocation_mode"`
	GroupBy                   string                          `json:"group_by"`
	OutputFormat              string                          `json:"output_format"`
	TraceId                   string                          `json:"trace_id"`
}

const k8sValidationURL = "https://app.getdigraph.com/api/validate/kubernetes"

func invokeDigraphKubernetesValidateAPI(kubernetesManifest kubernetes.ParsedKubernetesPlan, digraphAPIKey, mode, repository, ref, commitSHA, groupBy, outputFormat, traceId string, issueNumber int) (string, error) {
	requestBody := KubernetesConfigValidatorInput{
		KubernetesManifest: kubernetesManifest,
		Repository:         repository,
		Ref:                ref,
		InvocationMode:     mode,
		GroupBy:            groupBy,
		OutputFormat:       outputFormat,
		TraceId:            traceId,
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
		return "", &error_handling.RequestError{StatusCode: res.StatusCode, Err: fmt.Errorf("Error with API %s", body)}
	}

	return string(body), nil
}

func kubernetesRunCommand(cmd *cobra.Command) error {
	kubernetesManifest, _ := cmd.Flags().GetString("kubernetes-manifest")

	if len(kubernetesManifest) == 0 {
		return fmt.Errorf("must include kubernetes-manifest yaml")
	}

	digraphAPIKey, _ := cmd.Flags().GetString("api-key")
	repository, _ := cmd.Flags().GetString("repository")
	ref, _ := cmd.Flags().GetString("ref")

	issueNumber, _ := cmd.Flags().GetInt("issue-number")
	commitSHA, _ := cmd.Flags().GetString("commit-sha")

	groupBy, _ := cmd.Flags().GetString("group-by")
	outputFormat, _ := cmd.Flags().GetString("output-format")

	traceId, _ := cmd.Flags().GetString("traceId")

	if len(digraphAPIKey) == 0 {
		err := godotenv.Load(".env")

		if err != nil {
			return fmt.Errorf("must specify api-key as argument or set it within a .env file")
		}

		digraphAPIKey = os.Getenv("api-key")
	}

	parsedJSON, err := kubernetes.ParseKubernetesYAMLToJSON(kubernetesManifest)
	if err != nil {
		return fmt.Errorf("error parsing kubernetes manifest: %s", err)
	}

	mode := "cli"
	if len(commitSHA) > 0 || issueNumber > 0 {
		mode = "ci/cd"
	}

	output, err := invokeDigraphKubernetesValidateAPI(parsedJSON, digraphAPIKey, mode, repository, ref, commitSHA, groupBy, outputFormat, traceId, issueNumber)
	if err != nil {
		return err
	}

	if mode == "cli" || outputFormat == "json" {
		fmt.Fprintf(OutputWriter, "%s\n", output)
	}
	return nil
}

func validateKubernetes() *cobra.Command {
	k8sCmd := &cobra.Command{
		Use:       "kubernetes",
		Short:     "Validate kubernetes infra config changes",
		Long:      ``,
		ValidArgs: []string{"--", "-"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return kubernetesRunCommand(cmd)
		},
	}

	k8sCmd.Flags().String("kubernetes-manifest", "", "Kubernetes manifest YAML")
	return k8sCmd
}
