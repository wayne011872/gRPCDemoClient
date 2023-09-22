package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	pbh "github.com/wayne011872/gRPCDemoClient/proto/hello"
)

const (
	address     = "localhost:30000"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pbh.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := "Miles"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pbh.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}