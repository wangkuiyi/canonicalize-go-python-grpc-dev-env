package main

import (
	"context"
	"net"
	"testing"
	"time"

	pb "github.com/wangkuiyi/multi-stream-grpc/proto"
	"google.golang.org/grpc"
)

func TestRunJob(t *testing.T) {
	lis, e := net.Listen("tcp", ":0")
	if e != nil {
		t.Fatalf("failed to listen: %v", e)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSQLFlowServer(grpcServer, &sqlFlowServer{})
	go grpcServer.Serve(lis)

	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, e := grpc.Dial(lis.Addr().String(), opts...)
	if e != nil {
		t.Fatalf("fail to dial: %v", e)
	}
	defer conn.Close()
	client := pb.NewSQLFlowClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	j, e := client.Run(ctx, &pb.SQL{Sql: "SELECT *"})
	if j.Id != 100 {
		t.Errorf("Expecting 100, got %d", j.Id)
	}
}
