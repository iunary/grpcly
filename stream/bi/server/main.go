package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	pb "github.com/iunary/grpcly/stream/bi/proto"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50000, "grpc server port")
)

type Server struct {
	pb.UnimplementedReverserServer
}

func (s *Server) Broadcast(stream pb.Reverser_BroadcastServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Println(err.Error())
			return err
		}
		reverser := func(word string) (result string) {
			for _, v := range word {
				result = string(v) + result
			}
			return
		}
		log.Println("received ", msg.Word)
		stream.Send(&pb.ReverserResponse{
			Result: reverser(msg.Word),
		})
	}
}

func main() {
	flag.Parse()
	listner, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalln(err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterReverserServer(s, &Server{})
	log.Printf("server listening on localhost:%d", *port)
	if err := s.Serve(listner); err != nil {
		log.Fatalln(err.Error())
	}

}
