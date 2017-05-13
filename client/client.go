package client

import (
	"fmt"
	"net"
	"time"
)

// TCPClient struct
type TCPClient struct {
	Host string
	Port int
}

// Start TCPClient
func (c *TCPClient) Start() {
	message := "hello from client"
	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", c.Host, c.Port))
	if err != nil {
		panic(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	for {
		conn.Write([]byte(message))
		time.Sleep(time.Duration(1) * time.Second)
		reply := make([]byte, 1024)
		_, err = conn.Read(reply)
		if err != nil {
			panic(err)
		}
		fmt.Printf("received from server: [%s]\n", string(reply))
	}
}
