package cmd

import (
	"fmt"
	"os"

	"github.com/eimlav/adhocio/config"

	"github.com/spf13/cobra"
)

// CLI flags
var (
	Experimental bool
	ConfigPath   string
	RunRef       string
	RunRepo      string
	RunUser      string
)

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
	runCmd.Flags().StringVarP(&RunRef, "ref", "", "master", "Git ref of the checkout; can be branchName, for tag use refs/tags/tagName, commitId (SHA), etc (default: master)")
	runCmd.Flags().StringVarP(&RunRef, "repo", "", "", "The repository in Github to clone. (default: set by job)")
	runCmd.Flags().StringVarP(&RunUser, "user", "", "", "User/org in Github to clone the target repo from (default: set by job)")
	rootCmd.PersistentFlags().StringVarP(&ConfigPath, "filepath", "c", os.Getenv("HOME")+"/.adhocio.yaml", "Path to configuration file")
	rootCmd.AddCommand(configCmd, enableCmd, runCmd, versionCmd)
}

func initConfig() {
	err := config.GetConfig(ConfigPath)
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
