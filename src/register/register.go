package register

import (
	"github.com/pt-suzuki/grpc_template/infrastructure/connection"
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

	db, err := connection.ProvideDB("")
	if err != nil {
	}

	if err = db.Connect(); err != nil {
	}

	wire := provider.Wire(db.DB())
	pb.RegisterRockPaperScissorsServiceServer(server, wire.RockPaperScissorsProvider)

	return server
}
