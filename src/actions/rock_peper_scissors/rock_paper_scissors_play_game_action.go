package rock_peper_scissors

import (
	"context"
	"github.com/pt_suzuki/grpc_template/src/domains/rock_peper_scissors"
	"github.com/pt_suzuki/grpc_template/src/generate/pb"
	"github.com/pt_suzuki/grpc_template/src/generate/pb/models/rock_paper_scissors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type prayGameAction struct {
	service    rock_peper_scissors.Service
	translator rock_peper_scissors.Translator
}

func PlayGameActionImpl(
	service rock_peper_scissors.Service,
	translator rock_peper_scissors.Translator,
) pb.RockPaperScissorsServiceServer {
	return &prayGameAction{
		service:    service,
		translator: translator,
	}
}

func (s *prayGameAction) PlayGammmm(_ context.Context, req *pb.PlayRequest) (*pb.PlayResponse, error) {
	if req.HandShapes == rock_paper_scissors.HandShapes_HAND_SHAPES_UNKNOWN {
		return nil, status.Errorf(codes.InvalidArgument, "Choose Rock, Paper, or Scissors.")
	}

	matchResult := s.service.PlayGame(req.HandShapes)

	return &pb.PlayResponse{
		MatchResult: matchResult,
	}, nil
}

func (s *prayGameAction) ReportMatchResults(_ context.Context, _ *pb.ReportRequest) (*pb.ReportResponse, error) {
	return &pb.ReportResponse{
		Report: s.service.ReportMatchResults(),
	}, nil
}
