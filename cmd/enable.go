package cmd

import (
	"fmt"
	"time"

	"github.com/eimlav/adhocio/constants"

	"github.com/eimlav/adhocio/utils"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable all jobs in a pipeline",
	Long:  `Enable all jobs in a pipeline`,
	Args:  cobra.ExactArgs(1),
	Run:   enableCmdHandler,
}

func enableCmdHandler(cmd *cobra.Command, args []string) {
	// Get config
	var jobPrefix string
	if Experimental {
		jobPrefix = viper.GetString("experimental_prefix")
	} else {
		jobPrefix = viper.GetString("adhoc_prefix")
	}
	domain := viper.GetString("jenkins_domain")
	jobs := viper.GetStringSlice("jobs")

	// Enable each job
	successCount := 0
	for jobIndex := 0; jobIndex < len(jobs); jobIndex++ {
		time.Sleep(constants.API_SLEEP_DELAY * time.Millisecond)
		jobURL := fmt.Sprintf("https://%s/job/%s-%s_%s/enable", domain, jobPrefix, args[0], jobs[jobIndex])
		var output map[string]interface{}
		fmt.Println("Executing request to " + jobURL)
		if err := utils.MakeAPIRequest(jobURL, "POST", output); err != nil {
			fmt.Printf("Error occured: %v\n", err)
			continue
		}
		successCount++
	}

	// Print results
	fmt.Printf("\n%d of %d jobs enabled successfully!", successCount, len(jobs))
}
