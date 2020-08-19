package main

import (
	"flag"
	"github.com/raylax/logd/agent"
	"github.com/raylax/logd/client"
	"github.com/raylax/logd/server"
	"io/ioutil"
	"log"
)

// 命令行参数
var (
	serverMode bool
	agentMode bool
	clientMode bool
	upstream string
	config string
	color bool
)

func init()  {
	flag.BoolVar(&serverMode, "server", false, "server mode")
	flag.BoolVar(&agentMode, "agent", false, "agent mode")
	flag.StringVar(&config, "config", "", "config file")

	flag.BoolVar(&clientMode, "client", false, "client mode")
	flag.BoolVar(&clientMode, "color", false, "color console (client mode only)")
	flag.StringVar(&upstream, "upstream", "", "upstream server (client mode only)")
	flag.Parse()
}


func main() {
	
	log.SetFlags(log.Lshortfile | log.LstdFlags)


	var configString string
	data, err := ioutil.ReadFile(config)
	if err == nil {
		configString = string(data)
	}

	if clientMode {
		config := client.ParseConfig(configString)
		config.Color = color
		if upstream != "" {
			config.Upstream = upstream
		}
		client.NewClient(config).Start()
		return
	}

	if serverMode {
		config := server.ParseConfig(configString)
		server.NewServer(config).Start()
		return
	}

	if agentMode && configString != "" {
		config := agent.ParseConfig(configString)
		agent.NewAgent(config).Start()
		return
	}

	flag.Usage()
}
