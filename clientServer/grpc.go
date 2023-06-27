package main

import (
	"context"
	"fmt"
	"net"

	pb "Danial-Movahed.github.io/clientServerGrpc"
	"google.golang.org/grpc"
)

var port = 50051

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) RequestHoliday(ctx context.Context, in *pb.HolidayRequest) (*pb.Holiday, error) {
	fmt.Printf("Received: %s %s %s\n", in.GetDay(), in.GetMonth(), in.GetYear())
	return &pb.Holiday{Name: "test", Description: fmt.Sprintf("%s %s %s", in.GetDay(), in.GetMonth(), in.GetYear())}, nil
}

func StartGrpcServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	fmt.Printf("server listening at %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
	}
}
