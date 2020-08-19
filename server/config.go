package server

import (
	"gopkg.in/yaml.v2"
	"log"
)

type config struct {
	Listen listen `yaml:"listen"` // 监听
}

type listen struct {
	Agent string `yaml:"agent"` // agent监听
	Client string `yaml:"client"`// client监听
}

func ParseConfig(str string) *config {
	config := &config{}
	err := yaml.Unmarshal([]byte(str), config)
	if err != nil {
		log.Println(err)
	}
	return config
}
