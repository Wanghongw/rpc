package server

import (
	"github.com/golang/protobuf/protoc-gen-go/grpc"
	"log"
	"net"
)

const (
	port = "41005"
)

type Data struct{}

func main() {
	//起服务
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	inf.RegisterDataServer(s, &Data{})
	s.Serve(lis)

	log.Println("grpc server in: %s", port)
}
