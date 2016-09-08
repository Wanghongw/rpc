package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"os"
	"rpc/logger"
	"rpc/thrift-go/gen-go/hello/demo"
	"time"
)

func main() {
	logger.Init()

	transport, err := thrift.NewTSocket(net.JoinHostPort("127.0.0.1", "19090"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}
	transport.SetTimeout(time.Hour * 7)

	tF := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	pF := thrift.NewTBinaryProtocolFactoryDefault()
	client := demo.NewHelloThriftClientFactory(tF.GetTransport(transport), pF)

	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to 127.0.0.1:19090", " ", err)
		os.Exit(1)
	}
	defer transport.Close()

	goroutines := 3
	ch := make(chan int, goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			for range ch {
				_, err := client.SayHello(&demo.HelloRequest{})
				if err != nil {
					logger.ErrorStd("%v", err)
				}
			}
		}()
	}
	now := time.Now()
	for i := 0; i < 100; i++ {
		ch <- i
	}
	logger.DebugStd("%v", time.Since(now))
}
