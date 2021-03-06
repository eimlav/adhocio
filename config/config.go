package config

import (
	"errors"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

// GetConfig - Retrieves the configuration stored at the specified filepath
func GetConfig(filepath string) error {
	if filepath != "" {
		viper.SetConfigFile(filepath)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			return err
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".adhocio.yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if viper.GetString("jenkins_domain") == "" {
		return errors.New("No Jenkins domain found in config.")
	}

	if len(viper.GetStringSlice("jobs")) == 0 {
		return errors.New("No jobs found in config.")
	}

	if os.Getenv("JENKINS_USERNAME") == "" || os.Getenv("JENKINS_API_TOKEN") == "" {
		return errors.New("Ensure that JENKINS_USERNAME and JENKINS_API_TOKEN env vars are set.")
	}

	return nil
}
