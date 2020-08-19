package server

import (
	"github.com/raylax/logd/model"
	"github.com/raylax/logd/netio"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const (
	defaultAgentListen  = "0.0.0.0:9876"
	defaultClientListen = "0.0.0.0:6789"
	bufferSize          = 1024
)

type server struct {
	config  *config
	logChan chan *model.LogLine
	writerMap map[string]io.Writer
	mux sync.Mutex
}

func NewServer(config *config) *server {
	return &server{config: config}
}

func (server *server) Start() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)
	server.logChan = make(chan *model.LogLine, bufferSize)
	server.writerMap = make(map[string]io.Writer)
	agentListen := server.config.Listen.Agent
	if agentListen == "" {
		agentListen = defaultAgentListen
	}
	go listenTcp(agentListen, server.handleAgent)
	clientListen := server.config.Listen.Client
	if clientListen == "" {
		clientListen = defaultClientListen
	}
	go listenTcp(clientListen, server.handleClient)
	go server.listenChan()
	<-done
}

func (server *server) handleAgent(conn *net.TCPConn) {
	log.Printf("handle agent [%s]", conn.RemoteAddr().String())
	for {
		logLine, err := netio.ReadLog(conn)
		if err != nil {
			break
		}
		server.logChan <- logLine
	}
}

func (server *server) handleClient(conn *net.TCPConn) {
	server.mux.Lock()
	remoteAddr := conn.RemoteAddr()
	server.writerMap[remoteAddr.String()] = conn
	server.mux.Unlock()
}

func (server *server) listenChan()  {
	for {
		logLine := <-server.logChan
		server.mux.Lock()
		for k, writer := range server.writerMap {
			go func() {
				err := netio.WriteLog(writer, logLine)
				if err != nil {
					log.Printf("write to %s error, %v", k, err)
					server.removeClient(k)
				}
			}()
		}
		server.mux.Unlock()
	}
}

func (server *server) removeClient(key string)  {
	server.mux.Lock()
	delete(server.writerMap, key)
	server.mux.Unlock()
}

func listenTcp(address string, f func(conn *net.TCPConn)) {
	log.Printf("listening address [%s]", address)
	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		log.Fatalln(err)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Printf("listener accept error, %v", err)
		}
		go f(conn)
	}
}
