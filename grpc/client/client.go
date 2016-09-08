package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	gr "rpc/gorpc"
	"rpc/logger"
)

func main() {
	logger.Init()
	conn, err := grpc.Dial("localhost:8088", grpc.WithInsecure())
	if err != nil {
		logger.ErrorStd("%v", err)
	}
	defer conn.Close()

	c := gr.NewGreeterClient(conn)

	repl, err := c.SayHello(context.Background(), &gr.HelloRequest{Name: "yp"})
	if err != nil {
		logger.ErrorStd("%v", err)
	}
	logger.DebugStd("repl...%#v", repl)
}
