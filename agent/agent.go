package agent

import (
	"github.com/nxadm/tail"
	"github.com/raylax/logd/model"
	"github.com/raylax/logd/netio"
	"log"
	"net"
	"os"
	"time"
)

const bufferSize = 1024

type agent struct {
	config *config
	connected bool
}

func NewAgent(config *config) *agent {
	return &agent{config:config}
}

func (a *agent) Start()  {
	logs := make(chan *model.LogLine, bufferSize)
	server := a.config.Server
	log.Printf("connecting to upstream [%s]", a.config.Upstream)
	addr, err := net.ResolveTCPAddr("tcp", a.config.Upstream)
	if err != nil {
		log.Fatalf("%v", err)
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatalf("%v", err)
	}
	for _, logFile := range a.config.Logs {
		go a.tailFile(logs, server, logFile.File, logFile.Name)
	}
	a.connected = true
	for line := range logs {
		err = netio.WriteLog(conn, line)
		if err != nil {
			a.connected = false
			log.Printf("write to upstream error, %v", err)
			for {
				<-time.After(1 * time.Second)
				log.Printf("reconnecting to upstream [%s]", a.config.Upstream)
				conn, err = net.DialTCP("tcp", nil, addr)
				if err != nil {
					log.Printf("%v", err)
					continue
				}
				log.Printf("upstream reconnected [%s]", a.config.Upstream)
				a.connected = true
				break
			}
		}
	}
}

func (a *agent) tailFile(logs chan<- *model.LogLine, server string, file string, name string) {
	log.Printf("open file [%s]", file)
	stat, err := os.Stat(file)
	if err != nil {
		log.Printf("can not stat file [%s], error:%v", file, err)
	}
	logFile, err := tail.TailFile(file, tail.Config{Follow: true, Location:&tail.SeekInfo{
		Offset: stat.Size(),
		Whence: 0,
	}})
	if err != nil {
		log.Printf("can not read file [%s], error:%v", file, err)
	}
	for line := range logFile.Lines {
		if !a.connected {
			continue
		}
		logs <- &model.LogLine{
			Time:	line.Time.Unix(),
			Server: server,
			File:   name,
			Line:   line.Text,
		}
	}
}