package agent

import (
	"gopkg.in/yaml.v2"
	"log"
)

type config struct {
	// 上游服务器
	Upstream string `yaml:"upstream"`
	// 上报服务器名
	Server string `yaml:"server"`
	// 日志文件
	Logs []logFile `yaml:"logs"`
}

type logFile struct {
	// 文件
	File string `yaml:"file"`
	// 上报文件名，默认取file的文件名
	Name string `yaml:"name"`
}

func ParseConfig(str string) *config {
	config := &config{}
	err := yaml.Unmarshal([]byte(str), config)
	if err != nil {
		log.Println(err)
	}
	return config
}

