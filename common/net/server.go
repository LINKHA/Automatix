package net

import (
	"fmt"
	"net"
	"os"
)

type ServerConfig struct {
	Name string
	IP   string
	Port int
}

type Server struct {
	Name     string
	IP       string
	Port     int
	exitChan chan struct{}

	ConnMgr *ConnManager
}

func NewServer(config *ServerConfig) *Server {
	s := &Server{
		Name:     config.Name,
		IP:       config.IP,
		Port:     config.Port,
		exitChan: nil,

		ConnMgr: newConnManager(),
	}

	return s
}

func (s *Server) ListenTcpConn() {
	// TCP Server
	address := fmt.Sprintf("%s:%d", s.IP, s.Port)

	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		fmt.Println("Error resolving TCP address:", err)
		os.Exit(1)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("Error listening for TCP:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("TCP server is running on", tcpAddr)

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Error accepting TCP connection:", err)
				continue
			}

		}
	}()

	select {
	case <-s.exitChan:
		err := listener.Close()
		if err != nil {
			fmt.Println("listener close err: %v", err)
		}
	}
}
