package rock_peper_scissors

type Report struct {
	ID            string
	NumberOfGames uint8
	NumberOfWins  uint8
	MatchResults  RockPaperScissors
	UserId        string
}

func (r *Report) CalcGameCount() {
	r.NumberOfGames = r.NumberOfGames + 1

	for _, item := range r.MatchResults {
		if item.Result == WIN {
			r.NumberOfWins = r.NumberOfWins + 1
		}
	}
}
