package rock_peper_scissors

import (
	"context"
	"github.com/pt-suzuki/grpc_template/src/domains/rock_peper_scissors"
	"github.com/pt-suzuki/grpc_template/src/generate/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PlayGameResponder interface {
	PlayGameResponderInvoke(cxt context.Context, result *rock_peper_scissors.RockPaperScissor, err error) (*pb.PlayResponse, error)
}

type prayGameResponder struct {
	translator rock_peper_scissors.Translator
}

func PlayGameResponderImpl(translator rock_peper_scissors.Translator) PlayGameResponder {
	return &prayGameResponder{
		translator: translator,
	}
}

func (r *prayGameResponder) PlayGameResponderInvoke(_ context.Context, result *rock_peper_scissors.RockPaperScissor, err error) (*pb.PlayResponse, error) {
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unknown Error")
	}
	return &pb.PlayResponse{
		MatchResult: r.translator.ModelToPBMatchResult(result),
	}, nil
}
