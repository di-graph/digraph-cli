/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
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

func init() {
	validate.Flags().String("raw-output-plan", "", "Terminal output from terraform plan command")
	validate.Flags().String("output-path-plan", "", "Filepath for terminal output from terraform plan command")
	validate.Flags().String("json-path-plan", "", "Filepath to terraform plan JSON file")
	validate.Flags().String("json-output-plan", "", "JSON output from terraform plan command")

	validate.Flags().String("terraform-api-key", "", "Terraform API Key")
	validate.Flags().String("terraform-workspace", "", "Terraform workspace for associated plan")

	validate.PersistentFlags().String("api-key", "", "Digraph API Key")

	validate.PersistentFlags().String("repository", "", "Github repository")
	validate.PersistentFlags().String("ref", "", "Branch ref")
	validate.PersistentFlags().Int("issue-number", 0, "Pull Request Number")
	validate.PersistentFlags().String("commit-sha", "", "Commit SHA")

	validate.PersistentFlags().String("mode", "ci/cd", "Running mode- ci/cd or cli")

	rootCmd.AddCommand(validate)
	validate.AddCommand(validateKubernetes())
	validate.AddCommand(validateTerraform())
}
