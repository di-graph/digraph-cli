/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

var validate = &cobra.Command{
	Use:       "validate",
	Short:     "Validate infra config changes",
	Long:      ``,
	ValidArgs: []string{"--", "-"},
	RunE: func(cmd *cobra.Command, args []string) error {
		// for backwards compatability we run the terraform command when no sub command specified
		return terraformRunCommand(cmd)
	},
}

func FlagErrorFunc(cmd *cobra.Command, err error) error {
	if err == nil {
		return err
	}

	if strings.Contains(err.Error(), "unknown flag") {
		return &FlagParsingError{Err: err}
	}
	return err
}

func init() {
	// persistent flags are flags that are shared between all sub commands
	validate.PersistentFlags().String("api-key", "", "Digraph API Key")

	validate.PersistentFlags().String("raw-output-plan", "", "Terminal output from terraform plan command")
	validate.PersistentFlags().String("output-path-plan", "", "Filepath for terminal output from terraform plan command")
	validate.PersistentFlags().String("json-path-plan", "", "Filepath to terraform plan JSON file")
	validate.PersistentFlags().String("json-output-plan", "", "JSON output from terraform plan command")

	validate.PersistentFlags().String("terraform-api-key", "", "Terraform API Key")
	validate.PersistentFlags().String("terraform-workspace", "", "Terraform workspace for associated plan")

	validate.PersistentFlags().String("repository", "", "Github repository")
	validate.PersistentFlags().String("ref", "", "Branch ref")
	validate.PersistentFlags().Int("issue-number", 0, "Pull Request Number")
	validate.PersistentFlags().String("commit-sha", "", "Commit SHA")

	validate.PersistentFlags().String("group-by", "resource", "Group results based on resource or policy")
	validate.PersistentFlags().String("output-format", "terminal", "Format in which to output results")
	validate.SetFlagErrorFunc(FlagErrorFunc)
	validate.AddCommand(validateKubernetes())
	validate.AddCommand(validateTerraform())
}
