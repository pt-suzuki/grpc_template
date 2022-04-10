package provider

import "github.com/pt-suzuki/grpc_template/src/provider/pb_providers"

type Providers struct {
	RockPaperScissorsProvider pb_providers.RockPaperScissorsProvider
}

func ProvidersImpl(
	rockPaperScissorsProvider pb_providers.RockPaperScissorsProvider,
) *Providers {
	return &Providers{
		RockPaperScissorsProvider: rockPaperScissorsProvider,
	}
}
