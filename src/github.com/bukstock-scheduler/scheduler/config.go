package scheduler

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Tickers []string `yaml:"tickers"`
}

func (c *Config) parse() {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("Error while reading yaml file 'config.yaml' #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Printf("Error while parsing yaml file into struct #%v ", err)
	}
}
