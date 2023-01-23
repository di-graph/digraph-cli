/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"net/http"
	"os"

	"github.com/di-graph/digraph/internal/analytics"
	"github.com/spf13/cobra"
)

const VERSION = "v0.0.23"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "digraph",
	Version: VERSION,
	Short:   "Digraph CLI for policy configuration validation",
	Long:    ` Digraph CLI for policy configuration validation. To invoke, please call digraph validate terraform with the appropriate flags.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.SetVersionTemplate("Digraph {{.Version}}\n")

	analyticsWrapper := analytics.New()

	analyticsWrapper.SetClient(func() *http.Client {
		return http.DefaultClient
	})

	analyticsWrapper.SetVersion(VERSION)
	analyticsWrapper.SetCmdArgs(os.Args[1:])
	analyticsWrapper.SetTraceId()
	apiKey := analytics.GetAPIKey(os.Args[1:])
	analyticsWrapper.SetAPIKey(apiKey)
	traceId := analyticsWrapper.GetTraceId()
	rootCmd.PersistentFlags().String("traceId", traceId, "")
	rootCmd.PersistentFlags().MarkHidden("traceId")
	defer analyticsWrapper.Send()

	err := rootCmd.Execute()
	if err != nil {
		analyticsWrapper.AddError(err)
		analyticsWrapper.Send()
		_, ok := err.(*FlagParsingError)
		if ok {
			// we do not want to error out explicitly on flag parsing issues. Instead fail gracefully
			os.Exit(0)
		}
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.digraph.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
