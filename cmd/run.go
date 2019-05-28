package cmd

import (
	"fmt"

	"github.com/eimlav/adhocio/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a pipeline",
	Long:  `Run a pipeline`,
	Args:  cobra.ExactArgs(1),
	Run:   runCmdHandler,
}

func runCmdHandler(cmd *cobra.Command, args []string) {
	// Get config
	var jobPrefix string
	if Experimental {
		jobPrefix = viper.GetString("experimental_prefix")
	} else {
		jobPrefix = viper.GetString("adhoc_prefix")
	}
	domain := viper.GetString("jenkins_domain")
	jobs := viper.GetStringSlice("jobs")

	// Get parameters
	parameterString := ""
	if RunRef == "" {
		fmt.Printf("Invalid ref value passed")
		return
	}
	parameterString = fmt.Sprintf("?GITHUB_REF=%s", RunRef)

	if RunUser != "" {
		parameterString = fmt.Sprintf("%s&GITHUB_USER=%s", parameterString, RunUser)
	}

	if RunRepo != "" {
		parameterString = fmt.Sprintf("%s&GITHUB_REPO=%s", parameterString, RunRepo)
	}

	// Run adhoc init job with selected parameters
	jobURL := fmt.Sprintf("https://%s/job/%s-%s_%s/buildWithParameters%s", domain, jobPrefix, args[0], jobs[0], parameterString)

	var output map[string]interface{}
	fmt.Println("Executing request to " + jobURL)
	if err := utils.MakeAPIRequest(jobURL, "POST", output); err != nil {
		fmt.Printf("Error occured: %v\n", err)
		return
	}

	// Print results
	fmt.Print("Init job kicked off successfully!")
}
