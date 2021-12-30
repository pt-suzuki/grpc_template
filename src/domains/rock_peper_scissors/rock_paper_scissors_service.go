package rock_peper_scissors

import (
	"github.com/pt_suzuki/grpc_template/src/generate/pb/models/rock_paper_scissors"
	"math/rand"
)

type Service interface {
	PlayGame(handShapes rock_paper_scissors.HandShapes) *rock_paper_scissors.MatchResult
	ReportMatchResults() *rock_paper_scissors.Report
}

type service struct {
	numberOfGames int32
	numberOfWins  int32
	matchResults  []*rock_paper_scissors.MatchResult
	translator    Translator
}

func ServiceImpl(translator Translator) Service {
	return &service{
		numberOfGames: 0,
		numberOfWins:  0,
		matchResults:  make([]*rock_paper_scissors.MatchResult, 0),
		translator:    translator,
	}
}

func (s *service) PlayGame(handShape rock_paper_scissors.HandShapes) *rock_paper_scissors.MatchResult {
	// ランダムに１~３の数値を生成し相手の手とし、`HandShapes`のenumに変換しています。
	opponentHandShapes := s.translator.EncodeHandShapes(int32(rand.Intn(3) + 1))

	// ジャンケンの勝敗を決めています。
	result := s.playHandShape(handShape, opponentHandShapes)

	// 自動生成された型を元に対戦結果を生成
	matchResult := &rock_paper_scissors.MatchResult{
		YourHandShapes:     handShape,
		OpponentHandShapes: opponentHandShapes,
		Result:             result,
	}

	// 試合数を１増やし、プレイヤーが勝利した場合は勝利数も１増やします。
	s.addGameCount(matchResult)

	return matchResult
}

func (s *service) ReportMatchResults() *rock_paper_scissors.Report {
	return &rock_paper_scissors.Report{
		NumberOfGames: s.numberOfGames,
		NumberOfWins:  s.numberOfWins,
		MatchResults:  s.matchResults,
	}
}

func (s *service) playHandShape(handShape rock_paper_scissors.HandShapes, opponentHandShapes rock_paper_scissors.HandShapes) rock_paper_scissors.Result {
	if handShape == opponentHandShapes {
		return rock_paper_scissors.Result_DRAW
	} else if (handShape.Number()-opponentHandShapes.Number()+3)%3 == 1 {
		return rock_paper_scissors.Result_WIN
	}
	return rock_paper_scissors.Result_LOSE
}

func (s *service) addGameCount(matchResult *rock_paper_scissors.MatchResult) {
	s.numberOfGames = s.numberOfGames + 1
	if matchResult.Result == rock_paper_scissors.Result_WIN {
		s.numberOfWins = s.numberOfWins + 1
	}
	s.matchResults = append(s.matchResults, matchResult)
}
