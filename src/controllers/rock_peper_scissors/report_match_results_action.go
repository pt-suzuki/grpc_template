package rock_peper_scissors

import (
	"context"
	"github.com/pt-suzuki/grpc_template/src/domains/rock_peper_scissors"
	"github.com/pt-suzuki/grpc_template/src/generate/pb"
)

type ReportMatchResultsAction interface {
	ReportMatchResultsActionInvoke(cxt context.Context, _ *pb.ReportRequest) (*pb.ReportResponse, error)
}

type reportMatchResultsAction struct {
	useCase    rock_peper_scissors.ReportUseCase
	translator rock_peper_scissors.ReportTranslator
	responder  ReportMatchResultsResponder
}

func ReportMatchResultsActionImpl(
	useCase rock_peper_scissors.ReportUseCase,
	translator rock_peper_scissors.ReportTranslator,
	responder ReportMatchResultsResponder,
) ReportMatchResultsAction {
	return &reportMatchResultsAction{
		useCase:    useCase,
		translator: translator,
		responder:  responder,
	}
}

func (s *reportMatchResultsAction) ReportMatchResultsActionInvoke(cxt context.Context, req *pb.ReportRequest) (*pb.ReportResponse, error) {
	result, err := s.useCase.GetByUserId(req.UserId)
	return s.responder.ReportMatchResultsInvoke(cxt, result, err)
}
