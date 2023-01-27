package main

import (
	"context"
	"flag"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/iunary/grpcly/stream/ss/proto"
)

var (
	addr = flag.String("addr", "localhost:50000", "Grpc server address")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer conn.Close()

	c := pb.NewPassworderClient(conn)

	stream, err := c.Generate(context.Background(), &pb.PassworderRequest{
		Length: 8,
		Count:  6,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln(err.Error())
		}

		log.Println(msg.Password)
	}
}
