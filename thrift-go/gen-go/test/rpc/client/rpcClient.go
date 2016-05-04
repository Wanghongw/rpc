package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"os"
	"rpc/logger"
	"rpc/thrift-go/gen-go/test/rpc"
	"time"
)

func init() {
	logger.Init()
}
func main() {
	startTime := currentTimeMillis()
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort("127.0.0.1", "19090"))
	defer transport.Close()
	if err != nil {
		logger.ErrorStd("%v", err)
		os.Exit(1)
	}
	useTransport := transportFactory.GetTransport(transport)
	client := rpc.NewRpcServiceClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		logger.ErrorStd("Error opening socket to 127.0.0.1:19090 %v", err)
		os.Exit(1)
	}

	for i := 0; i < 1000000; i++ {
		paramMap := make(map[string]string)
		paramMap["name"] = "qinerg"
		paramMap["passwd"] = "123456"
		r1, e1 := client.FunCall(currentTimeMillis(), paramMap)
		logger.DebugStd("Call->%v:%v", r1, e1)
	}

	endTime := currentTimeMillis()
	logger.DebugStd("Program exit. time->共耗时:%vms", (endTime - startTime))
}

// 转换成毫秒
func currentTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}
