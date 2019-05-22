package cmd

import (
	"fmt"
	"os"

	"github.com/eimlav/adhocio/config"

	"github.com/spf13/cobra"
)

// Experimental flag value
var Experimental bool
var ConfigPath string
var Conf *config.Config

var rootCmd = &cobra.Command{
	Use:   "adhocio",
	Short: "CLI tool for running adhoc pipelines on Jenkins",
	Long: `Adhocio is a CLI for running adhoc pipelines on Jenkins.
https://github.com/eimlav/adhocio 
			`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello :) enter 'adhocio help' to get started")
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	enableCmd.Flags().BoolVarP(&Experimental, "experimental", "x", false, "Use experimental adhoc")
	runCmd.Flags().BoolVarP(&Experimental, "experimental", "x", false, "Use experimental adhoc")
	rootCmd.PersistentFlags().StringVarP(&ConfigPath, "filepath", "c", os.Getenv("HOME")+"/.adhocio.yaml", "Path to configuration file")
	rootCmd.AddCommand(runCmd, enableCmd, versionCmd)
}

func initConfig() {
	_, err := config.GetConfig(ConfigPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Execute - Runs CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
