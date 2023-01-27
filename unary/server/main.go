package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sort"

	pb "github.com/iunary/grpcly/unary/proto"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50000, "Grpc port")
)

type Server struct {
	pb.UnimplementedAnagramServer
}

func (s *Server) Check(ctx context.Context, in *pb.AnagramRequest) (*pb.AnagramResponse, error) {
	if len(in.Source) != len(in.Target) {
		return &pb.AnagramResponse{
			IsAnagram: false,
		}, nil
	}

	source := []byte(in.Source)
	sort.Slice(source, func(i, j int) bool {
		return source[i] < source[j]
	})

	target := []byte(in.Target)
	sort.Slice(target, func(i, j int) bool {
		return target[i] < target[j]
	})

	return &pb.AnagramResponse{
		IsAnagram: bytes.Equal(source, target),
	}, nil
}

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatal(err.Error())
	}
	s := grpc.NewServer()
	pb.RegisterAnagramServer(s, &Server{})
	log.Printf("server listening on localhost:%d", *port)
	if err := s.Serve(listener); err != nil {
		log.Fatal(err.Error())
	}
}
