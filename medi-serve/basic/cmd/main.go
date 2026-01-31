package main

import (
	"fmt"
	"google.golang.org/grpc"
	_ "medi-biz/common/init"
	"medi-biz/medi-serve/handler/service"
	pb "medi-biz/proto/medi" //注意这个路径
	"net"
)

func main() {
	// 地址
	addr := "127.0.0.1:50051"
	// 1.监听
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("监听异常:%s\n", err)
	}
	fmt.Printf("监听端口：%s\n", addr)
	// 2.实例化gRPC
	s := grpc.NewServer()
	// 3.在gRPC上注册微服务
	pb.RegisterMediServer(s, &service.Service{})
	// 4.启动服务端
	s.Serve(listener)
}
