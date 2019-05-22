package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable all jobs in a pipeline",
	Long:  `Enable all jobs in a pipeline`,
	Run:   enableCmdHandler,
}

func enableCmdHandler(cmd *cobra.Command, args []string) {
	fmt.Println("Not implemented")
}
