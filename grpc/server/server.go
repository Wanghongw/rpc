package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	gr "rpc/gorpc"
	"rpc/logger"
)

const (
	port = "8088"
)

type ServerImpl struct{}

func (this *ServerImpl) SayHello(ctx context.Context, in *gr.HelloRequest) (*gr.HelloReply, error) {
	return &gr.HelloReply{Message: "你好: " + in.Name}, nil
}

func main() {
	logger.Init()
	//起服务
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.ErrorStd("%v", err)
	}
	logger.DebugStd("grpc server in: %s", port)
	s := grpc.NewServer()
	gr.RegisterGreeterServer(s, &ServerImpl{})
	s.Serve(lis)
}
