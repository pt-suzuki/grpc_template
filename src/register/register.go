package register

import (
	server2 "github.com/pt-suzuki/grpc_template/infrastructure/server"
	"github.com/pt-suzuki/grpc_template/src/generate/pb"
	"github.com/pt-suzuki/grpc_template/src/provider"
	"google.golang.org/grpc"
)

type Register interface {
	GetRegisterServer() *grpc.Server
}

type register struct{}

func RegisterImpl() Register {
	return &register{}
}

func (r *register) GetRegisterServer() *grpc.Server {
	server := server2.CreateGrpc()

	wire := provider.Wire()

	// 自動生成された関数に、サーバと実際に処理を行うメソッドを実装したハンドラを設定します。
	// protoファイルで定義した`RockPaperScissorsService`に対応しています。
	pb.RegisterRockPaperScissorsServiceServer(server, wire.RockPaperScissorsProvider)

	return server
}
