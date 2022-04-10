package rock_peper_scissors

import (
	"context"
	"github.com/pt-suzuki/grpc_template/src/generate/pb"
	"github.com/pt-suzuki/grpc_template/src/generate/pb/models/rock_paper_scissors"
)

type PlayGameResponder interface {
	PlayGameResponderInvoke(cxt context.Context, matchResult *rock_paper_scissors.MatchResult) (*pb.PlayResponse, error)
}

type prayGameResponder struct {
}

func PlayGameResponderImpl() PlayGameResponder {
	return &prayGameResponder{}
}

func (r *prayGameResponder) PlayGameResponderInvoke(cxt context.Context, matchResult *rock_paper_scissors.MatchResult) (*pb.PlayResponse, error) {
	return &pb.PlayResponse{
		MatchResult: matchResult,
	}, nil
}
