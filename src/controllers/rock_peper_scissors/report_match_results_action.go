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
	service    rock_peper_scissors.Service
	translator rock_peper_scissors.Translator
	responder  ReportMatchResultsResponder
}

func ReportMatchResultsActionImpl(
	service rock_peper_scissors.Service,
	translator rock_peper_scissors.Translator,
	responder ReportMatchResultsResponder,
) ReportMatchResultsAction {
	return &reportMatchResultsAction{
		service:    service,
		translator: translator,
		responder:  responder,
	}
}

func (s *reportMatchResultsAction) ReportMatchResultsActionInvoke(cxt context.Context, _ *pb.ReportRequest) (*pb.ReportResponse, error) {
	return s.responder.ReportMatchResultsInvoke(cxt, s.service.ReportMatchResults())
}
