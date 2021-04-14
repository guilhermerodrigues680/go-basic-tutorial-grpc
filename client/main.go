package main

import (
	"app/pb"
	"context"
	"flag"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name used to verify the hostname returned by the TLS handshake")
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewFullNameBuilderClient(conn)

	fullNameRequest := pb.FullNameRequest{
		FirstName: "Joao",
		LastName:  "Luz",
	}

	log.Printf("Looking for full name within %s %s", fullNameRequest.FirstName, fullNameRequest.LastName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fullName, err := client.GetFullName(ctx, &fullNameRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(fullName)

	//
	stream0, err := client.GetNames(ctx, fullName)
	if err != nil {
		log.Fatal(err)
	}

	for {
		name, err := stream0.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetNames(_) = _, %v", client, err)
		}
		log.Println(name)
	}

	//
	stream2, err2 := client.GetFullNameWithNames(ctx)
	if err2 != nil {
		log.Fatal(err)
	}

	name := pb.Name{
		Name: "Jose",
	}

	for i := 0; i < 3; i++ {
		if err := stream2.Send(&name); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream2, name.Name, err)
		}
		log.Printf("%v.Send(%v)", stream2, name.Name)
	}

	reply, err := stream2.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream2, err, nil)
	}
	log.Printf("Route summary: %v", reply)

	//
	{
		stream3, err3 := client.GetNamesStream(ctx)
		if err3 != nil {
			log.Fatal(err)
		}
		// for i := 0; i < 3; i++ {
		// 	if err := stream3.Send(&name); err != nil {
		// 		log.Fatalf("%v.Send(%v) = %v", stream3, name.Name, err)
		// 	}
		// 	log.Printf("%v.Send(%v)", stream3, name.Name)
		// }

		waitc := make(chan struct{})
		go func() {
			for {
				in, err := stream3.Recv()
				if err == io.EOF {
					// read done.
					close(waitc)
					break
				}
				if err != nil {
					log.Fatalf("Failed to receive a note : %v", err)
				}
				log.Printf("Got message %s", in.Name)
			}
		}()
		for i := 0; i < 5; i++ {
			if err := stream3.Send(&name); err != nil {
				log.Fatalf("Failed to send a name: %v", err)
			}
		}
		stream3.CloseSend()
		<-waitc

	}
}
