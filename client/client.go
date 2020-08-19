package client

import (
	"github.com/raylax/logd/console"
	"github.com/raylax/logd/netio"
	"log"
	"net"
)

type client struct {
	config *config
}

func NewClient(config *config) *client {
	return &client{config:config}
}

func (c *client) Start()  {
	if c.config.Upstream == "" {
		print("upstream must be not empty")
		return
	}
	log.Printf("connecting to upstream [%s]", c.config.Upstream)
	addr, err := net.ResolveTCPAddr("tcp", c.config.Upstream)
	if err != nil {
		log.Fatalf("%v", err)
		return
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("upstream connected")
	for {
		logLine, err := netio.ReadLog(conn)
		if err != nil {
			log.Fatalln(err)
		}
		console.Println(logLine)
	}
}

