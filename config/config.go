package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	JenkinsDomain      string   `yaml:"jenkins_domain"`
	AdhocPrefix        string   `yaml:"adhoc_prefix"`
	ExperimentalPrefix string   `yaml:"experimental_prefix"`
	JobPrefix          string   `yaml:"job_prefix"`
	Jobs               []string `yaml:"jobs"`
}

// GetConfig - Retrieves the configuration stored at the specified filepath
func GetConfig(filepath string) (*Config, error) {
	configFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	config := &Config{}

	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
