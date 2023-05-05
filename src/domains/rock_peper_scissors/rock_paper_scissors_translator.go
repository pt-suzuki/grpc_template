package rock_peper_scissors

import (
	"github.com/pt-suzuki/grpc_template/src/generate/pb/models/rock_paper_scissors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Translator interface {
	RequestHandShapes(n rock_paper_scissors.HandShapes) HandShapes
	ModelToPBMatchResult(content *RockPaperScissor) *rock_paper_scissors.MatchResult
	ModelsToPBMatchResults(list RockPaperScissors) []*rock_paper_scissors.MatchResult
	ModelToMapper(m *RockPaperScissor) *RockPaperScissorGorm
	ModelToUpdateInterface(m *RockPaperScissor) map[string]interface{}
	MappersToModels(ms []*RockPaperScissorGorm) RockPaperScissors
	MapperToModel(m *RockPaperScissorGorm) *RockPaperScissor
}

type translator struct {
}

func TranslatorImpl() Translator {
	return &translator{}
}

func (t *translator) RequestHandShapes(h rock_paper_scissors.HandShapes) HandShapes {
	switch h {
	case 1:
		return ROCK
	case 2:
		return PAPER
	case 3:
		return SCISSORS
	default:
		return HAND_SHAPES_UNKNOWN
	}
}

func (t *translator) ModelsToPBMatchResults(ms RockPaperScissors) []*rock_paper_scissors.MatchResult {
	result := make([]*rock_paper_scissors.MatchResult, 0, len(ms))
	for _, item := range ms {
		result = append(result, t.ModelToPBMatchResult(item))
	}
	return result
}

func (t *translator) ModelToPBMatchResult(m *RockPaperScissor) *rock_paper_scissors.MatchResult {
	return &rock_paper_scissors.MatchResult{
		YourHandShapes:     rock_paper_scissors.HandShapes(m.YourHandShapes),
		OpponentHandShapes: rock_paper_scissors.HandShapes(m.OpponentHandShapes),
		Result:             rock_paper_scissors.Result(m.Result),
		CreateTime:         timestamppb.New(m.CreateTime),
	}
}

func (t *translator) ModelToMapper(m *RockPaperScissor) *RockPaperScissorGorm {
	return &RockPaperScissorGorm{
		ID:                 m.ID,
		YourHandShapes:     uint32(m.YourHandShapes),
		OpponentHandShapes: uint32(m.OpponentHandShapes),
		Result:             uint32(m.Result),
		CreateTime:         m.CreateTime,
		UserID:             m.UserID,
	}
}

func (t *translator) ModelToUpdateInterface(m *RockPaperScissor) map[string]interface{} {
	return map[string]interface{}{
		"YourHandShapes":     uint32(m.YourHandShapes),
		"OpponentHandShapes": uint32(m.OpponentHandShapes),
		"Result":             uint32(m.Result),
		"UserID":             m.UserID,
	}
}

func (t *translator) MappersToModels(ms []*RockPaperScissorGorm) RockPaperScissors {
	result := make(RockPaperScissors, 0, len(ms))
	for _, item := range ms {
		result = append(result, t.MapperToModel(item))
	}
	return result
}

func (t *translator) MapperToModel(m *RockPaperScissorGorm) *RockPaperScissor {
	return &RockPaperScissor{
		ID:                 m.ID,
		YourHandShapes:     HandShapes(m.YourHandShapes),
		OpponentHandShapes: HandShapes(m.OpponentHandShapes),
		Result:             Result(m.Result),
		CreateTime:         m.CreateTime,
		UserID:             m.UserID,
	}
}
