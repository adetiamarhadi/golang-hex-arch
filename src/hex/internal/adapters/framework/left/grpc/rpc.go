package grpc

import (
	"context"

	"github.com/adetiamarhadi/golang-hex-arch/internal/adapters/framework/left/grpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetAdd ...
func (a Adapter) GetAdd(ctx context.Context, req *pb.OperationParameter) (*pb.Answer, error) {

	var err error

	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := a.api.GetAdd(req.A, req.B)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &pb.Answer{Value: answer}

	return ans, nil
}

// GetSubtract ...
func (a Adapter) GetSubtract(ctx context.Context, req *pb.OperationParameter) (*pb.Answer, error) {

	var err error

	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := a.api.GetSubtract(req.A, req.B)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &pb.Answer{Value: answer}

	return ans, nil
}

// GetMultiply ...
func (a Adapter) GetMultiply(ctx context.Context, req *pb.OperationParameter) (*pb.Answer, error) {

	var err error

	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := a.api.GetMultiply(req.A, req.B)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &pb.Answer{Value: answer}

	return ans, nil
}

// GetDivide ...
func (a Adapter) GetDivide(ctx context.Context, req *pb.OperationParameter) (*pb.Answer, error) {

	var err error

	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := a.api.GetDivide(req.A, req.B)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &pb.Answer{Value: answer}

	return ans, nil
}
