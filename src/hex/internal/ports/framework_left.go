package ports

import (
	"context"

	"github.com/adetiamarhadi/golang-hex-arch/internal/adapters/framework/left/grpc/pb"
)

// GRPCPort ...
type GRPCPort interface {
	Run()
	GetAdd(ctx context.Context, req *pb.OperationParameter) (*pb.Answer, error)
	GetSubtract(ctx context.Context, req *pb.OperationParameter) (*pb.Answer, error)
	GetMultiply(ctx context.Context, req *pb.OperationParameter) (*pb.Answer, error)
	GetDivide(ctx context.Context, req *pb.OperationParameter) (*pb.Answer, error)
}
