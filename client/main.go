// Package main implements a simple gRPC client that demonstrates how to use gRPC-Go libraries
// to perform unary, client streaming, server streaming and full duplex RPCs.
//
// It interacts with the route guide service whose definition can be found in proto/route_guide.proto.
package main

import (
	"flag"
	"fmt"

	pb "github.com/letterj/grpcrest/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "testdata/ca.pem", "The file containning the CA root cert file")
	serverAddr         = flag.String("server_addr", "127.0.0.1:8443", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name to verify  hostname returned by TLS handshake")
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		var sn string
		if *serverHostOverride != "" {
			sn = *serverHostOverride
		}
		var creds credentials.TransportAuthenticator
		if *caFile != "" {
			var err error
			creds, err = credentials.NewClientTLSFromFile(*caFile, sn)
			if err != nil {
				grpclog.Fatalf("Failed to create TLS credentials %v", err)
			}
		} else {
			creds = credentials.NewClientTLSFromCert(nil, sn)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewGRPCRestApiClient(conn)

	// Create
	cdata, err := client.CreateWidget(context.Background(), &pb.CreateWRequest{Data: "test"})
	if err != nil {
		grpclog.Fatalf("Fatal gRPC error: %v", err)
	}
	fmt.Println(cdata.Result)

	// Show
	sdata, err := client.ShowWidget(context.Background(), &pb.ShowWRequest{Id: "test"})
	if err != nil {
		grpclog.Fatalf("Fatal gRPC error: %v", err)
	}
	fmt.Println(sdata.Result)

	// Delete
	ddata, err := client.DeleteWidget(context.Background(), &pb.DeleteWRequest{Id: "test"})
	if err != nil {
		grpclog.Fatalf("Fatal gRPC error: %v", err)
	}
	fmt.Println(ddata.Result)

	// Update
	udata, err := client.UpdateWidget(context.Background(), &pb.UpdateWRequest{Data: "test"})
	if err != nil {
		grpclog.Fatalf("Fatal gRPC error: %v", err)
	}
	fmt.Println(udata.Result)

	// List
	ldata, err := client.ListWidget(context.Background(), &pb.ListWRequest{})
	if err != nil {
		grpclog.Fatalf("Fatal gRPC error: %v", err)
	}
	fmt.Println(ldata.Result)

}
