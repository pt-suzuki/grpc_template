package provider

import (
	"github.com/pt_suzuki/grpc_template/src/actions/rock_peper_scissors"
	"github.com/pt_suzuki/grpc_template/src/generate/pb"
)

type ActionProvider struct {
	RockPaperScissorsPlayGameAction pb.RockPaperScissorsServiceServer
}

func GetActionProvider() ActionProvider {
	sp := GetServiceProvider()

	ap := ActionProvider{}
	ap.RockPaperScissorsPlayGameAction = rock_peper_scissors.PlayGameActionImpl(sp.RockPaperScissorsService, sp.RockPaperScissorsTranslator)

	return ap
}
