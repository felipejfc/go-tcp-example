package server

import (
	"fmt"
	"io"
	"net"
)

// TCPServer struct
type TCPServer struct {
	Bind string
	Port int
}

// Start TCPServer
func (s *TCPServer) Start() {
	fmt.Printf("started tcp echo server...\n")
	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Bind, s.Port))
	defer ln.Close()
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go func(conn net.Conn) {
			_, err := io.Copy(conn, conn)
			defer conn.Close()
			if err != nil {
				panic(err)
			}
		}(conn)
	}
}
