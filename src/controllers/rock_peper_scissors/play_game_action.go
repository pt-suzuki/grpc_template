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
	useCase    rock_peper_scissors.UseCase
	translator rock_peper_scissors.Translator
	responder  PlayGameResponder
}

func PlayGameActionImpl(
	useCase rock_peper_scissors.UseCase,
	translator rock_peper_scissors.Translator,
	responder PlayGameResponder,
) PlayGameAction {
	return &prayGameAction{
		useCase:    useCase,
		translator: translator,
		responder:  responder,
	}
}

func (a *prayGameAction) PlayGameActionInvoke(cxt context.Context, req *pb.PlayRequest) (*pb.PlayResponse, error) {
	if req.HandShapes == rock_paper_scissors.HandShapes_HAND_SHAPES_UNKNOWN {
		return nil, status.Errorf(codes.InvalidArgument, "Choose Rock, Paper, or Scissors.")
	}

	matchResult, err := a.useCase.PlayGame(a.translator.RequestHandShapes(req.HandShapes))

	return a.responder.PlayGameResponderInvoke(cxt, matchResult, err)
}
