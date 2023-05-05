package rock_peper_scissors

import "github.com/golang/protobuf/ptypes/timestamp"

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
	YourHandShapes     HandShapes
	OpponentHandShapes HandShapes
	Result             Result
	CreateTime         timestamp.Timestamp
}

func (m *RockPaperScissor) PlayHandShape() Result {
	if m.YourHandShapes == m.OpponentHandShapes {
		return DRAW
	} else if (m.YourHandShapes-m.OpponentHandShapes+3)%3 == 1 {
		return WIN
	}
	return LOSE
}

type RockPaperScissors []*RockPaperScissor
