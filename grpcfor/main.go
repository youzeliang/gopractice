package main

import (
	"context"
	"fmt"
	"github.com/youzeliang/gopractice/grpcfor/example"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

var times = 0

type service struct {
}

func (s *service) HelloWorld(ctx context.Context, in *example.HelloRequest) (*example.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	times++
	return &example.HelloReply{Message: "Hello 您好！" + in.Name+" 您是今天第"+strconv.Itoa(times)+"个访客"}, nil
}

func main() {
	conn, err := net.Listen("tcp", "127.0.0.1:9999")
	fmt.Println(conn, err)

	server := grpc.NewServer()


	example.RegisterHelloServer(server, &service{})
	//读取请求并响应
	err = server.Serve(conn)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}