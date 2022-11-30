package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func main() {
	// 开启服务端口监听
	s, _ := net.Listen("tcp", ":9999")
	grpcServer := grpc.NewServer()

	fmt.Printf("start grpc-server %s-%s on %d success...\n", "127.0.0.1", "v1", 9999)
	// 启动rpc server
	err := grpcServer.Serve(s)
	if err != nil {
		fmt.Println("start grpc-server failed ", err)
	}
}
