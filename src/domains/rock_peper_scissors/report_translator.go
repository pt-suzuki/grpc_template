package rock_peper_scissors

import "github.com/pt-suzuki/grpc_template/src/generate/pb/models/rock_paper_scissors"

type ReportTranslator interface {
	ContentToPB(content *Report) *rock_paper_scissors.Report
}

type reportTranslator struct {
	translator Translator
}

func ReportTranslatorImpl(translator Translator) ReportTranslator {
	return &reportTranslator{
		translator: translator,
	}
}

func (r reportTranslator) ContentToPB(content *Report) *rock_paper_scissors.Report {
	return &rock_paper_scissors.Report{
		NumberOfWins:  int32(content.NumberOfWins),
		NumberOfGames: int32(content.NumberOfGames),
		MatchResults:  r.translator.ModelsToPBMatchResults(content.MatchResults),
		UserId:        content.UserId,
	}
}
