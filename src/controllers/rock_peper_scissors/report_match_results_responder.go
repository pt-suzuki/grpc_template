package rock_peper_scissors

import (
	"context"
	"github.com/pt-suzuki/grpc_template/src/generate/pb"
	"github.com/pt-suzuki/grpc_template/src/generate/pb/models/rock_paper_scissors"
)

type ReportMatchResultsResponder interface {
	ReportMatchResultsInvoke(cxt context.Context, report *rock_paper_scissors.Report) (*pb.ReportResponse, error)
}

type reportMatchResultsResponder struct {
}

func ReportMatchResultsResponderImpl() ReportMatchResultsResponder {
	return &reportMatchResultsResponder{}
}

func (r *reportMatchResultsResponder) ReportMatchResultsInvoke(cxt context.Context, report *rock_paper_scissors.Report) (*pb.ReportResponse, error) {
	return &pb.ReportResponse{
		Report: report,
	}, nil
}
