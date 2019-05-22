package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Experimental flag value
var Experimental bool

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
	enableCmd.Flags().BoolVarP(&Experimental, "experimental", "x", false, "Use experimental adhoc")
	runCmd.Flags().BoolVarP(&Experimental, "experimental", "x", false, "Use experimental adhoc")
	rootCmd.AddCommand(runCmd, enableCmd, versionCmd)
}

// Execute - Runs CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
