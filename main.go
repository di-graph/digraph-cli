/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/di-graph/digraph/cmd"
)

func main() {
	rootCmd := cmd.RootCmd()
	cmd.Execute(rootCmd)
}
