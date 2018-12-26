package scheduler

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// Config Configuration
type Config struct {
	TimeoutMillis int      `yaml:"timeout-millis"`
	Tickers       []string `yaml:"tickers"`
}

// Parse load yaml into memory
func (c *Config) Parse() {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("Error while reading yaml file 'config.yaml' #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Printf("Error while parsing yaml file into struct #%v ", err)
	}
}
