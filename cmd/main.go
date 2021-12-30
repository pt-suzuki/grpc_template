package main

import (
	"fmt"
	"github.com/pt_suzuki/grpc_template/src/router"
	"log"
	"net"
)

func main() {
	// 起動するポート番号を指定しています。
	port := 50051
	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := router.RouterImpl().GetRouter()

	// サーバーを起動
	server.Serve(listenPort)
}
