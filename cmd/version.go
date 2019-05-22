package cmd

import (
	"fmt"

	"github.com/eimlav/adhocio/constants"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print current version of adhocio",
	Long:  `Print current version of adhocio`,
	Run:   versionCmdHandler,
}

func versionCmdHandler(cmd *cobra.Command, args []string) {
	fmt.Println("adhocio " + constants.VERSION)
}
