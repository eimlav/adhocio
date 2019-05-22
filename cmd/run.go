package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a pipeline",
	Long:  `Run a pipeline`,
	Run:   runCmdHandler,
}

func runCmdHandler(cmd *cobra.Command, args []string) {
	fmt.Println("Not implemented")
}
