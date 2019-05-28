package cmd

import (
	"fmt"
	"os"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: fmt.Sprintf("Display config (default: %s/.adhocio.yaml", os.Getenv("HOME")),
	Long:  `Display config`,
	Run:   configCmdHandler,
}

func configCmdHandler(cmd *cobra.Command, args []string) {
	// Get config
	experimentalPrefix := viper.GetString("experimental_prefix")
	adhocPrefix := viper.GetString("adhoc_prefix")
	jenkinsDomain := viper.GetString("jenkins_domain")
	jobs := viper.GetStringSlice("jobs")

	// Print out config
	config := struct {
		ExperimentalPrefix string
		AdhocPrefix        string
		JenkinsDomain      string
		Jobs               []string
	}{
		experimentalPrefix,
		adhocPrefix,
		jenkinsDomain,
		jobs,
	}
	configTemplate, err := template.ParseFiles("templates/config.txt")
	if err != nil {
		fmt.Printf("Error occured: %v\n", err)
		return
	}
	err = configTemplate.Execute(os.Stdout, config)
	if err != nil {
		fmt.Printf("Error occured: %v\n", err)
		return
	}
}
