//go:build wireinject
// +build wireinject

package provider

import (
	"github.com/google/wire"
	rock_peper_scissors2 "github.com/pt-suzuki/grpc_template/src/controllers/rock_peper_scissors"
	"github.com/pt-suzuki/grpc_template/src/domains/rock_peper_scissors"
	"github.com/pt-suzuki/grpc_template/src/provider/pb_provider"
	"gorm.io/gorm"
)

func Wire(db *gorm.DB) *Providers {
	wire.Build(
		rock_peper_scissors.TranslatorImpl,
		rock_peper_scissors.RepositoryImpl,
		rock_peper_scissors.UseCaseImpl,
		rock_peper_scissors.ReportTranslatorImpl,
		rock_peper_scissors.ReportUseCaseImpl,
		rock_peper_scissors2.PlayGameResponderImpl,
		rock_peper_scissors2.PlayGameActionImpl,
		rock_peper_scissors2.ReportMatchResultsResponderImpl,
		rock_peper_scissors2.ReportMatchResultsActionImpl,
		pb_provider.RockPaperScissorsProviderImpl,
		ProvidersImpl,
	)
	return nil
}
