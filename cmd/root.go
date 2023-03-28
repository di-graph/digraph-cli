/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/di-graph/digraph/internal/analytics"
	"github.com/di-graph/digraph/internal/error_handling"
	"github.com/spf13/cobra"
)

const VERSION = "v0.0.35"

var OutputWriter io.Writer = os.Stdout

func RootCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "digraph",
		Version: VERSION,
		Short:   "Digraph CLI for policy configuration validation",
		Long:    ` Digraph CLI for policy configuration validation. To invoke, please call digraph validate terraform with the appropriate flags.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(rootCmd *cobra.Command) error {
	rootCmd.SetOut(OutputWriter)
	rootCmd.SetErr(OutputWriter)
	rootCmd.SetVersionTemplate("Digraph {{.Version}}\n")
	rootCmd.AddCommand(validate)

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
			fmt.Println(error_handling.FormatError(err.Error()))
			os.Exit(0)
		}
		reqErr, ok := err.(*error_handling.RequestError)
		if ok {
			// we want to handle certain status codes specifically
			if reqErr.StatusCode == http.StatusUnauthorized {
				fmt.Println(error_handling.FormatError("Unauthorized request- please check API key"))
				os.Exit(1)
			}
			return err
		}
		return err
	}
	return nil
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
