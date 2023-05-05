package report

import "github.com/pt-suzuki/grpc_template/src/domains/rock_peper_scissors"

type Report struct {
	NumberOfGames uint8
	NumberOfWins  uint8
	MatchResults  MatchResults
}

func (r *Report) AddGameCount(matchResult *MatchResult) {
	r.NumberOfGames = r.NumberOfGames + 1
	if matchResult.Result == rock_peper_scissors.WIN {
		r.NumberOfWins = r.NumberOfWins + 1
	}
	r.MatchResults = append(r.MatchResults, matchResult)
}
