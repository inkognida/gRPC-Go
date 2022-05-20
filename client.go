package main

import (
	"context"
	"flag"
	"grpcadder/api/proto/adderpb"
	"log"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatalf("Not enought args")
	}

	x, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	y, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := adderpb.NewAdderClient(conn)
	res, err := c.Add(context.Background(), &adderpb.AddRequest{X: int32(x), Y: int32(y)})
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Println(res.GetR())
}
