//go:build wireinject
// +build wireinject

package provider

import (
	"github.com/google/wire"
	rock_peper_scissors2 "github.com/pt-suzuki/grpc_template/src/controllers/rock_peper_scissors"
	"github.com/pt-suzuki/grpc_template/src/domains/rock_peper_scissors"
	"github.com/pt-suzuki/grpc_template/src/provider/pb_provider"
)

func Wire() *Providers {
	wire.Build(
		rock_peper_scissors.TranslatorImpl,
		rock_peper_scissors.ServiceImpl,
		rock_peper_scissors2.PlayGameResponderImpl,
		rock_peper_scissors2.PlayGameActionImpl,
		rock_peper_scissors2.ReportMatchResultsResponderImpl,
		rock_peper_scissors2.ReportMatchResultsActionImpl,
		pb_provider.RockPaperScissorsProviderImpl,
		ProvidersImpl,
	)
	return nil
}
