package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"os"
	"rpc/logger"
	"rpc/thrift-go/gen-go/hello/demo"
	"time"
)

type RpcServiceImpl struct {
}

func (this *RpcServiceImpl) SayHello(helloReq *demo.HelloRequest) (*demo.HelloReply, error) {
	return &demo.HelloReply{Message: "你好: " + helloReq.Name}, nil
}

func main() {
	logger.Init()
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	serverTransport, err := thrift.NewTServerSocketTimeout("localhost:19090", time.Hour*8)
	if err != nil {
		logger.Error("%v", err)
		os.Exit(1)
	}

	handler := &RpcServiceImpl{}
	processor := demo.NewHelloThriftProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	logger.DebugStd("thrift server in: %s", "19090")
	logger.ErrorStd("%v", server.Serve())
}
