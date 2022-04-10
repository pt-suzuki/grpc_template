package provider

import "github.com/pt-suzuki/grpc_template/src/provider/pb_provider"

type Providers struct {
	RockPaperScissorsProvider pb_provider.RockPaperScissorsProvider
}

func ProvidersImpl(
	rockPaperScissorsProvider pb_provider.RockPaperScissorsProvider,
) *Providers {
	return &Providers{
		RockPaperScissorsProvider: rockPaperScissorsProvider,
	}
}
