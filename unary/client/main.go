package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/iunary/grpcly/unary/proto"
)

var (
	addr   = flag.String("addr", "localhost:50000", "Grpc server address")
	source = flag.String("source", "", "")
	target = flag.String("target", "", "")
)

func main() {
	flag.Parse()

	if len(*source) == 0 || len(*target) == 0 {
		log.Fatalln("source and target must not be empty")
	}

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer conn.Close()

	c := pb.NewAnagramClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	r, err := c.Check(ctx, &pb.AnagramRequest{
		Source: *source,
		Target: *target,
	})

	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Printf("%s and %s are anagram [%t]", *source, *target, r.GetIsAnagram())
}
