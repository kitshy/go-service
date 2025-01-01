package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/kitshy/go-service/config"
	"github.com/kitshy/go-service/service/rpc"
)

func main() {

	fmt.Println("running grpc server ...")
	var f = flag.String("c", "config.yml", "config path")
	flag.Parse()
	conf, err := config.New(*f)
	if err != nil {
		fmt.Println("failed to load config:", err)
		return
	}
	port, err := strconv.Atoi(conf.Server.Port)
	if err != nil {
		fmt.Println("failed to parse server port:", err)
	}
	grpcServerCfg := &rpc.RpcServerConfig{
		GrpcHostname: conf.Server.Host,
		GrpcPort:     port,
	}

	rpcServer, err := rpc.NewRpcServer(grpcServerCfg)
	if err != nil {
		fmt.Println("failed to init rpc server:", err)
	}
	err = rpcServer.Start()
	if err != nil {
		fmt.Println("failed to start rpc server:", err)
		return
	}

	fmt.Println("grpc server started ")

	<-(chan struct{})(nil)

}

// 原生 rpc 服务
//func main() {
//
//	var f = flag.String("c", "config.yml", "config path")
//	flag.Parse()
//	conf, err := config.New(*f)
//	if err != nil {
//		panic(err)
//	}
//
//	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(nil))
//	defer grpcServer.GracefulStop()
//
//	listen, err := net.Listen("tcp", ":"+conf.Server.Port)
//	if err != nil {
//		fmt.Println("net listen failed err", err)
//		panic(err)
//	}
//
//	reflection.Register(grpcServer)
//
//	fmt.Println("grpc server start success:", conf.Server.Port)
//
//	if err := grpcServer.Serve(listen); err != nil {
//		fmt.Println("grpc server start failed err", err)
//		panic(err)
//	}
//
//}
