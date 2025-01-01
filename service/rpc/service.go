package rpc

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"sync/atomic"

	"github.com/kitshy/go-service/protobuf/wallet"
)

const MaxRecvMessageSize = 1024 * 1024 * 300

type RpcServerConfig struct {
	GrpcHostname string
	GrpcPort     int
}

type RpcServer struct {
	*RpcServerConfig
	wallet.UnimplementedWalletServiceServer
	stopped atomic.Bool
}

func (s *RpcServer) Stop() error {
	s.stopped.Store(true)
	return nil
}

func (s *RpcServer) Stopped() bool {
	return s.stopped.Load()
}

func NewRpcServer(config *RpcServerConfig) (*RpcServer, error) {
	return &RpcServer{RpcServerConfig: config}, nil
}

func (s *RpcServer) Start() error {
	go func(s *RpcServer) {

		address := fmt.Sprintf("%s:%d", s.RpcServerConfig.GrpcHostname, s.RpcServerConfig.GrpcPort)
		fmt.Println("start rpc service", "address", address)

		lis, err := net.Listen("tcp", address)
		if err != nil {
			fmt.Println("Could not start rpc service ,failed to listen")
		}

		opt := grpc.MaxRecvMsgSize(MaxRecvMessageSize)
		gs := grpc.NewServer(opt, grpc.ChainUnaryInterceptor(nil))

		reflection.Register(gs)

		wallet.RegisterWalletServiceServer(gs, s)
		fmt.Println("grpc service info", "port", s.GrpcPort, "address", lis.Addr().String())

		if err := gs.Serve(lis); err != nil {
			fmt.Println("grpc service error", "err", err)
		}

	}(s)
	return nil
}
