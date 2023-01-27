package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	pb "github.com/iunary/grpcly/stream/cs/proto"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50000, "Grpc port")
)

type Server struct {
	pb.UnimplementedReverserServer
}

func (s *Server) ReverseString(stream pb.Reverser_ReverseStringServer) error {
	reverser := func(word string) (result string) {
		for _, v := range word {
			result = string(v) + result
		}
		return
	}

	var result []string
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.ReverserResponse{
				Result: result,
			})
		}

		if err != nil {
			log.Println(err.Error())
			return err
		}

		log.Println("received", msg.Word)
		result = append(result, reverser(msg.Word))

	}

}

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalln(err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterReverserServer(s, &Server{})
	log.Printf("server listening on localhost:%d", *port)
	if err := s.Serve(listener); err != nil {
		log.Fatal(err.Error())
	}
}
