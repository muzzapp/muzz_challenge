package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"muzz_challenge/cmd/internal/types"
	explore "muzz_challenge/pkg/proto"
)

type Server struct {
	explore.UnimplementedExploreServiceServer
	repo repo
}

type repo interface {
	CountLikedYou(userId string) (int, error)
	ListNewLikedYou(
		userId string,
		paginationToken *string,
	) ([]*types.User, *string, error)
	PutDecision(deciderId string, decision *types.Decision) error
	ListLikedYou(userId string, pageSize int, paginationToken *string) ([]*types.User, *string, error)
}

func New(repo repo) *Server {
	return &Server{
		repo: repo,
	}
}

func (s *Server) ListLikedYou(ctx context.Context, in *explore.ListLikedYouRequest) (*explore.ListLikedYouResponse, error) {
	// TODO - implement me
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (s *Server) ListNewLikedYou(ctx context.Context, in *explore.ListLikedYouRequest) (*explore.ListLikedYouResponse, error) {
	// TODO - implement me
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (s *Server) PutDecision(ctx context.Context, in *explore.PutDecisionRequest) (*explore.PutDecisionResponse, error) {
	// TODO - implement me
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (s *Server) CountLikedYou(ctx context.Context, in *explore.CountLikedYouRequest) (*explore.CountLikedYouResponse, error) {
	// TODO - implement me
	return nil, status.Error(codes.Unimplemented, "not implemented")
}
