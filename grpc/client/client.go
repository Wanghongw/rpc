package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	gr "rpc/gorpc"
	"rpc/logger"
	"time"
)

func main() {
	logger.Init()
	conn, err := grpc.Dial("localhost:8088", grpc.WithInsecure())
	if err != nil {
		logger.ErrorStd("%v", err)
	}
	defer conn.Close()

	c := gr.NewGreeterClient(conn)

	goroutines := 100
	ch := make(chan int, goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			for range ch {
				_, err = c.SayHello(context.Background(), &gr.HelloRequest{Name: "yp"})
				if err != nil {
					logger.ErrorStd("%v", err)
				}
			}
		}()
	}
	now := time.Now()
	for i := 0; i < 1000000; i++ {
		ch <- i
	}
	logger.DebugStd("%v", time.Since(now))
}
