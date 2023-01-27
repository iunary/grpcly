package main

import (
	"context"
	"flag"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/iunary/grpcly/stream/bi/proto"
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
	c := pb.NewReverserClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream, err := c.Broadcast(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}

	done := make(chan struct{})
	words := []string{"foobar", "tarbar"}

	go func() {
		for _, word := range words {
			if err := stream.Send(&pb.ReverserRequest{
				Word: word,
			}); err != nil {
				log.Fatalln(err.Error())
			}
		}
	}()

	go func() {

		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}

			if err != nil {
				log.Fatalln(err.Error())
				close(done)
				return
			}

			log.Println("received ", msg.Result)
		}
	}()

	<-done
}
