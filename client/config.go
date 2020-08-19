package client

import (
	"gopkg.in/yaml.v2"
	"log"
)

type config struct {

	Upstream string `yaml:"upstream"`

	Color bool `yaml:"color"`

}

func ParseConfig(str string) *config {
	config := &config{}
	err := yaml.Unmarshal([]byte(str), config)
	if err != nil {
		log.Println(err)
	}
	return config
}
