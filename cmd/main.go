package main

import (
	"app/pb"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type appFullNameServer struct {
	pb.UnimplementedFullNameBuilderServer
}

// Recebe um nome e sobrenome e retorna um nome completo
func (s *appFullNameServer) GetFullName(context.Context, *pb.FullNameRequest) (*pb.FullName, error) {
	log.Println("GetFullName")
	return &pb.FullName{
		Name: "Oii",
	}, nil
}

// Recebe um nome completo e retorna um stream com varios nomes
func (s *appFullNameServer) GetNames(fullName *pb.FullName, stream pb.FullNameBuilder_GetNamesServer) error {
	for i := 0; i < 4; i++ {
		if err := stream.Send(&pb.Name{Name: "ola"}); err != nil {
			return err
		}
	}
	return nil
}

// Recebe varios nomes e retorna um nome completo
func (s *appFullNameServer) GetFullNameWithNames(stream pb.FullNameBuilder_GetFullNameWithNamesServer) error {
	for {
		name, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.FullName{
				Name: "Full name",
			})
		}
		if err != nil {
			return err
		}
		log.Println(name)
	}
	// return status.Errorf(codes.Unimplemented, "method GetFullNameWithNames not implemented")
}

// Recebe varios nomes e retorna varios nomes
func (s *appFullNameServer) GetNamesStream(stream pb.FullNameBuilder_GetNamesStreamServer) error {
	for {
		name, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		//key := serialize(in.Location)
		name.Name = name.Name + "sss"

		for i := 0; i < 2; i++ {
			if err := stream.Send(name); err != nil {
				return err
			}
		}
	}
	// return status.Errorf(codes.Unimplemented, "method GetNamesStream not implemented")
}

// func (s *appFullNameServer) mustEmbedUnimplementedFullNameBuilderServer() {}

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 10000, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	s := &appFullNameServer{}
	pb.RegisterFullNameBuilderServer(grpcServer, s)
	log.Println("Listening on: ", *port)
	grpcServer.Serve(lis)
}
