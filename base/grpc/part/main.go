package main

import (
	"log"
	"net"
	"net/rpc"
)

func main() {

	rpc.RegisterName("HelloService", new(HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error")
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("listen error")
	}

	rpc.ServeConn(conn)
}

type HelloService struct {
}

func (p *HelloService) Hello() error {

	return nil

}
