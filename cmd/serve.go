/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"Golabi-boilerplate/packages/bootstrap"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Bootstrap the application",
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.Serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
