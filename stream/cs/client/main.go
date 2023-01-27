package main

import (
	"context"
	"flag"
	"log"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/iunary/grpcly/stream/cs/proto"
)

var (
	addr = flag.String("addr", "localhost:50000", "Grpc server addr")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer conn.Close()

	c := pb.NewReverserClient(conn)
	stream, err := c.ReverseString(context.Background())
	if err != nil {
		log.Fatalln(err.Error())
	}

	words := []string{"foobar", "tarbar", "tarbaz"}
	for _, word := range words {
		err := stream.Send(&pb.ReverserRequest{
			Word: word,
		})
		if err != nil {
			log.Fatalln(err.Error())
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(strings.Join(res.Result, ","))

}
