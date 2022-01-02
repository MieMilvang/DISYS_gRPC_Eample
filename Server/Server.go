package main

import (
	"context"
	"log"
	"net"

	pb "github.com/MieMilvang/DISYS_gRPC_Eample/Proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type Server struct{
	pb.UnimplementedServiceServer
}

func (s *Server) Increment(ctx context.Context, in *pb.IncrementValue) (*pb.Return, error){
	log.Printf("The requested incremented value: %v  - from client %d", in.GetValue(), in.GetClientid())
	return &pb.Return{IncrementedValue: in.GetValue()}, nil
}

func main(){
	lis, err := net.Listen("tcp", port)
	if err !=nil{
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterServiceServer(s, &Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err!= nil{
		log.Fatalf("failed to serve: %v", err)
	}
}