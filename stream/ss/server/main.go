package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strings"

	pb "github.com/iunary/grpcly/stream/ss/proto"
	"google.golang.org/grpc"
)

var (
	port           = flag.Int("port", 50000, "Grpc server port")
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

type Server struct {
	pb.UnimplementedPassworderServer
}

func (s *Server) Generate(in *pb.PassworderRequest, stream pb.Passworder_GenerateServer) error {
	for i := 0; i < int(in.Count); i++ {
		if err := stream.Send(&pb.PassworderResponse{
			Password: generatePassword(int(in.Length), 1, 1, 1),
		}); err != nil {
			log.Fatalln(err.Error())
		}
	}
	return nil
}

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalln(err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterPassworderServer(s, &Server{})
	log.Printf("server listening on localhost:%d", *port)
	if err := s.Serve(listener); err != nil {
		log.Fatal(err.Error())
	}
}

func generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
	var password strings.Builder

	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}
