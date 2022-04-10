package rock_peper_scissors

import (
	"context"
	"github.com/pt-suzuki/grpc_template/src/domains/rock_peper_scissors"
	"github.com/pt-suzuki/grpc_template/src/generate/pb"
	"github.com/pt-suzuki/grpc_template/src/generate/pb/models/rock_paper_scissors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PlayGameAction interface {
	PlayGameActionInvoke(cxt context.Context, req *pb.PlayRequest) (*pb.PlayResponse, error)
}

type prayGameAction struct {
	service    rock_peper_scissors.Service
	translator rock_peper_scissors.Translator
	responder  PlayGameResponder
}

func PlayGameActionImpl(
	service rock_peper_scissors.Service,
	translator rock_peper_scissors.Translator,
	responder PlayGameResponder,
) PlayGameAction {
	return &prayGameAction{
		service:    service,
		translator: translator,
		responder:  responder,
	}
}

func (s *prayGameAction) PlayGameActionInvoke(cxt context.Context, req *pb.PlayRequest) (*pb.PlayResponse, error) {
	if req.HandShapes == rock_paper_scissors.HandShapes_HAND_SHAPES_UNKNOWN {
		return nil, status.Errorf(codes.InvalidArgument, "Choose Rock, Paper, or Scissors.")
	}

	matchResult := s.service.PlayGame(req.HandShapes)

	return s.responder.PlayGameResponderInvoke(cxt, matchResult)
}
