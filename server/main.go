package main

import (
	"flag"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	"golang.org/x/net/context"

	pb "github.com/letterj/grpcrest/proto"
)

var (
	tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "kingdom/server1.pem", "The TLS cert file")
	keyFile  = flag.String("key_file", "kingdom/server1.key", "The TLS key file")
	port     = flag.Int("port", 8443, "The server port")
)

type widgetServer struct {
	dataVault []string
}

// Add a widget
func (s *widgetServer) CreateWidget(ctx context.Context, r *pb.CreateWRequest) (*pb.CreateWResponse, error) {
	output := `{"status": "add OK"}`
	return &pb.CreateWResponse{Result: output}, nil
}

// Show a widget
func (s *widgetServer) ShowWidget(ctx context.Context, r *pb.ShowWRequest) (*pb.ShowWResponse, error) {
	output := `{"status": "show OK"}`
	return &pb.ShowWResponse{Result: output}, nil
}

// Delete a widget
func (s *widgetServer) DeleteWidget(ctx context.Context, r *pb.DeleteWRequest) (*pb.DeleteWResponse, error) {
	output := `{"status": "delete OK"}`
	return &pb.DeleteWResponse{Result: output}, nil
}

// Update a widget
func (s *widgetServer) UpdateWidget(ctx context.Context, r *pb.UpdateWRequest) (*pb.UpdateWResponse, error) {
	output := `{"status": "update OK"}`
	return &pb.UpdateWResponse{Result: output}, nil
}

// List a widget
func (s *widgetServer) ListWidget(ctx context.Context, r *pb.ListWRequest) (*pb.ListWResponse, error) {
	output := `{"status": "List OK"}`
	return &pb.ListWResponse{Result: output}, nil
}

// Create a new server
func newServer() *widgetServer {
	s := new(widgetServer)
	seed1 := `{"id": "1", "name": "Widget 1"}`
	seed2 := `{"id": "2", "name": "Widget 2"}`
	s.dataVault = []string{seed1, seed2}
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			grpclog.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterGRPCRestApiServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
