package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/brotherlogic/gobuildmaster/proto"
)

func (s *Server) Claim(ctx context.Context, req *pb.ClaimRequest) (*pb.ClaimResponse, error) {
	if s.claimed == req.GetJobName() {
		return nil, status.Errorf(codes.FailedPrecondition, "Already claimed")
	}
	return &pb.ClaimResponse{}, nil
}
