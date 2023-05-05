package rock_peper_scissors

import (
	"math/rand"
	"time"
)

type HandShapes uint8

const (
	HAND_SHAPES_UNKNOWN HandShapes = iota + 1
	ROCK
	PAPER
	SCISSORS
)

type Result uint8

const (
	RESULT_UNKNOWN Result = iota + 1
	WIN
	LOSE
	DRAW
)

type RockPaperScissor struct {
	ID                 string
	YourHandShapes     HandShapes
	OpponentHandShapes HandShapes
	Result             Result
	CreateTime         time.Time
	UserID             string
}

func (m *RockPaperScissor) PlayHandShape() {
	// 自動生成された型を元に対戦結果を生成
	m.OpponentHandShapes = HandShapes(rand.Intn(3) + 1)
	if m.YourHandShapes == m.OpponentHandShapes {
		m.Result = DRAW
		return
	} else if (m.YourHandShapes-m.OpponentHandShapes+3)%3 == 1 {
		m.Result = WIN
		return
	}
	m.Result = LOSE
}

type RockPaperScissors []*RockPaperScissor
