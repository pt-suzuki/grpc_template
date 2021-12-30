package server

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func CreateGrpc() *grpc.Server {
	// gRPCサーバーの生成
	server := grpc.NewServer()

	// サーバーリフレクションを有効にしています。
	// 有効にすることでシリアライズせずとも後述する`grpc_cli`で動作確認ができるようになります。
	reflection.Register(server)

	return server
}
