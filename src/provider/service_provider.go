package provider

import "github.com/pt_suzuki/grpc_template/src/domains/rock_peper_scissors"

type ServiceProvider struct {
	RockPaperScissorsService    rock_peper_scissors.Service
	RockPaperScissorsTranslator rock_peper_scissors.Translator
}

func GetServiceProvider() ServiceProvider {
	sp := ServiceProvider{}

	sp.RockPaperScissorsTranslator = rock_peper_scissors.TranslatorImpl()
	sp.RockPaperScissorsService = rock_peper_scissors.ServiceImpl(sp.RockPaperScissorsTranslator)

	return sp
}
