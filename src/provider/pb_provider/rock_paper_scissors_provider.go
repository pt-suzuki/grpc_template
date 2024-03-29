package pb_provider

import (
	"context"
	"github.com/pt-suzuki/grpc_template/src/controllers/rock_peper_scissors"
	"github.com/pt-suzuki/grpc_template/src/generate/pb"
)

type RockPaperScissorsProvider interface {
	pb.RockPaperScissorsServiceServer
}

type rockPaperScissorsProvider struct {
	playAction               rock_peper_scissors.PlayGameAction
	reportMatchResultsAction rock_peper_scissors.ReportMatchResultsAction
}

func RockPaperScissorsProviderImpl(
	playGameAction rock_peper_scissors.PlayGameAction,
	reportMatchResultsAction rock_peper_scissors.ReportMatchResultsAction,
) RockPaperScissorsProvider {
	return &rockPaperScissorsProvider{
		playGameAction,
		reportMatchResultsAction,
	}
}

func (h *rockPaperScissorsProvider) PlayGame(cxt context.Context, req *pb.PlayRequest) (*pb.PlayResponse, error) {
	return h.playAction.PlayGameActionInvoke(cxt, req)
}

func (h *rockPaperScissorsProvider) ReportMatchResults(cxt context.Context, req *pb.ReportRequest) (*pb.ReportResponse, error) {
	return h.reportMatchResultsAction.ReportMatchResultsActionInvoke(cxt, req)
}
