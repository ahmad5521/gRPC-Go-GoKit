package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"elm.sa/ahmasiri/rpc-server/endpoints"
	pb "elm.sa/ahmasiri/rpc-server/pb"
	"elm.sa/ahmasiri/rpc-server/service"
	transport "elm.sa/ahmasiri/rpc-server/transports"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"google.golang.org/grpc"
)

func main() {
	var logger log.Logger
	logger = log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	addservice := service.NewService(logger)
	addendpoint := endpoints.MakeEndpoints(addservice)
	grpcServer := transport.NewGRPCServer(addendpoint, logger)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grpcListener, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	go func() {
		baseServer := grpc.NewServer()
		pb.RegisterMathServiceServer(baseServer, grpcServer)
		level.Info(logger).Log("msg", "Server started successfully 🚀")
		baseServer.Serve(grpcListener)
	}()

	level.Error(logger).Log("exit", <-errs)
}
