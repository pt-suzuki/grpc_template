package rock_peper_scissors

import (
	"github.com/pt_suzuki/grpc_template/src/generate/pb/models/rock_paper_scissors"
)

type Translator interface {
	EncodeHandShapes(n int32) rock_paper_scissors.HandShapes
}

type translator struct {
}

func TranslatorImpl() Translator {
	return &translator{}
}

func (t *translator) EncodeHandShapes(n int32) rock_paper_scissors.HandShapes {
	switch n {
	case 1:
		return rock_paper_scissors.HandShapes_ROCK
	case 2:
		return rock_paper_scissors.HandShapes_PAPER
	case 3:
		return rock_paper_scissors.HandShapes_SCISSORS
	default:
		return rock_paper_scissors.HandShapes_HAND_SHAPES_UNKNOWN
	}
}
