package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string

var rootCmd = &cobra.Command{
	Use:   "Golabi-boilerplate",
	Short: "A brief description of your application",
	Long: `
	In Order to run the app use: 
	go run .\main.go serve 

	In Order to run the app with custom env file use: 
	go run .\main.go serve --config yourFileName
	`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is .env)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Config file must be env
	viper.SetConfigType("env")

	// Config file must be in your project directory
	viper.AddConfigPath(".")
	if configFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configFile)
	} else {
		// Use default env file
		viper.SetConfigName(".env")
	}

	// read in environment variables that match
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	// Print Loaded env file name
	fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
}
