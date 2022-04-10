package router

import (
	server2 "github.com/pt-suzuki/grpc_template/infrastructure/server"
	"github.com/pt-suzuki/grpc_template/src/generate/pb"
	"github.com/pt-suzuki/grpc_template/src/provider"
	"google.golang.org/grpc"
)

type Router interface {
	GetRouter() *grpc.Server
}

type router struct{}

func RouterImpl() Router {
	return &router{}
}

func (r *router) GetRouter() *grpc.Server {
	server := server2.CreateGrpc()

	actions := provider.GetActionProvider()

	// 自動生成された関数に、サーバと実際に処理を行うメソッドを実装したハンドラを設定します。
	// protoファイルで定義した`RockPaperScissorsService`に対応しています。
	pb.RegisterRockPaperScissorsServiceServer(server, actions.RockPaperScissorsPlayGameAction)

	return server
}
