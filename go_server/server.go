// server.go
package main

import (
	"context"
	"log"
	"net"

	pb "../common/streamlit_and_go" // 生成されたGoコードのパッケージパスに合わせて変更

	"google.golang.org/grpc"
)

// サーバー実装（Greeterサービス）
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHelloメソッドの実装
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received request for name: %v", in.GetName())
	return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

func main() {
	// ポート50051で待ち受け
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// gRPCサーバーの生成
	s := grpc.NewServer()
	// Greeterサービスの登録
	pb.RegisterGreeterServer(s, &server{})

	log.Printf("Server is listening on %v", lis.Addr())
	// サーバーの起動
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
