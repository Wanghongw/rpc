package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"os"
	"rpc/logger"
	"rpc/thrift-go/gen-go/test/rpc"
)

const (
	NetworkAddr = "127.0.0.1:19090"
)

func init() {
	logger.Init()
}

type RpcServiceImpl struct {
}

func (this *RpcServiceImpl) FunCall(callTime int64, paramMap map[string]string) (r []string, err error) {
	logger.DebugStd("-->FunCall:%v:%v", callTime, paramMap)

	for k, v := range paramMap {
		r = append(r, k+v)
	}
	return
}
func main() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	serverTransport, err := thrift.NewTServerSocket(NetworkAddr)
	if err != nil {
		logger.ErrorStd("%v", err)
		os.Exit(1)
	}

	handler := &RpcServiceImpl{}
	processor := rpc.NewRpcServiceProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	logger.DebugStd("thrift server in %v", NetworkAddr)
	server.Serve()
}
