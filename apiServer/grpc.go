package main

import (
	"context"
	"fmt"
	"time"

	pb "Danial-Movahed.github.io/apiServerGrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = "localhost:50051"

func grpcClient(day string, month string, year string) *holiday {
	// Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("did not connect: %v\n", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.RequestHoliday(ctx, &pb.HolidayRequest{Day: day, Month: month, Year: year})
	if err != nil {
		fmt.Printf("could not greet: %v\n", err)
	}
	return &holiday{Name: r.GetName(), Description: r.GetDescription()}
}
