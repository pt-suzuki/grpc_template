package rock_peper_scissors

import (
	"context"
	"github.com/pt-suzuki/grpc_template/src/domains/rock_peper_scissors"
	"github.com/pt-suzuki/grpc_template/src/generate/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ReportMatchResultsResponder interface {
	ReportMatchResultsInvoke(cxt context.Context, report *rock_peper_scissors.Report, err error) (*pb.ReportResponse, error)
}

type reportMatchResultsResponder struct {
	translator rock_peper_scissors.ReportTranslator
}

func ReportMatchResultsResponderImpl(translator rock_peper_scissors.ReportTranslator) ReportMatchResultsResponder {
	return &reportMatchResultsResponder{
		translator,
	}
}

func (r *reportMatchResultsResponder) ReportMatchResultsInvoke(_ context.Context, report *rock_peper_scissors.Report, err error) (*pb.ReportResponse, error) {
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unknown Error")
	}
	return &pb.ReportResponse{
		Report: r.translator.ContentToPB(report),
	}, nil
}
