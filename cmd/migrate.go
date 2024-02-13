/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"Golabi-boilerplate/packages/db"

	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate the database models",
	Long:  `It first connect to db and then auto migrate all the models`,
	Run: func(cmd *cobra.Command, args []string) {
		db.MigrateDB()
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
