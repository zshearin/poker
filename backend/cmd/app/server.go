package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	kitlog "github.com/go-kit/log"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	ps "github.com/zshearin/poker/backend/pkg/poker-service"
	"google.golang.org/grpc"

	p_v1alpha1 "github.com/zshearin/poker/backend/api/v1alpha1"
)

var (
	grpcAddr = ":7123"
	httpAddr = ":8088"
)

func newServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Poker service server command",
		Run:   runServerCmd,
	}

	//TODO - ADD FLAGS

	return cmd
}

func runServerCmd(cmd *cobra.Command, args []string) {

	var logger kitlog.Logger

	logger = kitlog.NewJSONLogger(kitlog.NewSyncWriter(os.Stderr))
	logger = kitlog.With(logger, "timestamp", kitlog.DefaultTimestampUTC)

	//Service endpoints
	var (
		pokerServiceLogger = kitlog.With(logger, "service", "poker")
		pokerService       = ps.NewService(pokerServiceLogger)

		pokerServiceProxyEndpoints = ps.CreateEndpoints(pokerService, pokerServiceLogger)
	)

	// gRPC transport
	var (
		grpcLogger      = kitlog.With(logger, "component", "grpc")
		psGrpcTransport = ps.NewGrpcTransport(pokerServiceProxyEndpoints, grpcLogger)
	)

	grpcServer := grpc.NewServer()
	{
		p_v1alpha1.RegisterPokerAPIServer(grpcServer, psGrpcTransport)
	}

	grpcGateway := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}),
	)

	grpcOpts := []grpc.DialOption{grpc.WithInsecure()}
	err := p_v1alpha1.RegisterPokerAPIHandlerFromEndpoint(context.Background(), grpcGateway, grpcAddr, grpcOpts)
	if err != nil {
		log.Fatalf("cannot create gRPC to HTTP Gateway server endpoints: %s", err)
	}

	errs := make(chan error, 2)

	go func() {
		lis, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			logger.Log("transport", "grpc", "during", "listen", "err", err)
			os.Exit(1)
		}
		logger.Log("transport", "grpc", "address", grpcAddr, "msg", "listening")
		errs <- grpcServer.Serve(lis)
	}()

	go func() {
		logger.Log("transport", "http", "address", httpAddr, "msg", "listening")
		//add log msg
		errs <- http.ListenAndServe(httpAddr, grpcGateway)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)
}