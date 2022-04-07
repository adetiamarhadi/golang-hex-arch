package grpc

import (
	"log"
	"net"

	"github.com/adetiamarhadi/golang-hex-arch/internal/adapters/framework/left/grpc/pb"
	"github.com/adetiamarhadi/golang-hex-arch/internal/ports"
	"google.golang.org/grpc"
)

// Adapter ...
type Adapter struct {
	api ports.APIPort
}

// NewAdapter ...
func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

// Run ...
func (a Adapter) Run() {

	var err error

	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}

	arithmeticServiceServer := a

	grpcServer := grpc.NewServer()

	pb.RegisterArithmeticServiceServer(grpcServer, arithmeticServiceServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to server gRPC server over port 9000: %v", err)
	}
}
