package main

//go:generate protoc -I ../proto ../proto/sqlflow.proto --go_out=plugins=grpc:../proto

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/wangkuiyi/canonicalize-go-python-grpc-dev-env/proto"
	"google.golang.org/grpc"
)

type sqlFlowServer struct {
}

func (s *sqlFlowServer) Run(ctx context.Context, sql *pb.SQL) (*pb.Job, error) {
	return &pb.Job{Id: 100}, nil
}

func main() {
	port := flag.Int("port", 10000, "The server port")
	flag.Parse()

	lis, e := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if e != nil {
		log.Fatalf("failed to listen: %v", e)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterSQLFlowServer(grpcServer, &sqlFlowServer{})
	grpcServer.Serve(lis)
}
