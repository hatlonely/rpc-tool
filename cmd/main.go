package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/hatlonely/go-kit/bind"
	"github.com/hatlonely/go-kit/config"
	"github.com/hatlonely/go-kit/flag"
	"github.com/hatlonely/go-kit/logger"
	"github.com/hatlonely/go-kit/refx"
	"github.com/hatlonely/go-kit/rpcx"

	"github.com/hatlonely/rpc-tool/api/gen/go/api"
	"github.com/hatlonely/rpc-tool/internal/service"
)

var Version string

type Options struct {
	flag.Options

	GrpcGateway rpcx.GrpcGatewayOptions
	Service     service.Options

	Logger struct {
		Info logger.Options
		Grpc logger.Options
	}
}

func main() {
	var options Options
	refx.Must(flag.Struct(&options, refx.WithCamelName()))
	refx.Must(flag.Parse(flag.WithJsonVal()))
	if options.Help {
		fmt.Println(flag.Usage())
		return
	}
	if options.Version {
		fmt.Println(Version)
		return
	}

	if options.ConfigPath == "" {
		options.ConfigPath = "config/base.json"
	}
	cfg, err := config.NewConfigWithBaseFile(options.ConfigPath, refx.WithCamelName())
	refx.Must(err)

	refx.Must(bind.Bind(&options, []bind.Getter{flag.Instance(), bind.NewEnvGetter(bind.WithEnvPrefix("RPC_TOOL")), cfg}, refx.WithCamelName()))

	grpcLog, err := logger.NewLoggerWithOptions(&options.Logger.Grpc, refx.WithCamelName())
	refx.Must(err)
	infoLog, err := logger.NewLoggerWithOptions(&options.Logger.Info, refx.WithCamelName())
	refx.Must(err)
	infoLog.With("options", options).Info("init config success")
	cfg.SetLogger(infoLog)

	refx.Must(cfg.Watch())
	defer cfg.Stop()

	svc, err := service.NewToolServiceWithOptions(&options.Service)
	refx.Must(err)

	grpcGateway, err := rpcx.NewGrpcGatewayWithOptions(&options.GrpcGateway, refx.WithCamelName())
	refx.Must(err)
	grpcGateway.SetLogger(infoLog, grpcLog)

	api.RegisterToolServiceServer(grpcGateway.GRPCServer(), svc)
	refx.Must(grpcGateway.RegisterServiceHandlerFunc(api.RegisterToolServiceHandlerFromEndpoint))
	grpcGateway.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	grpcGateway.Stop()
	infoLog.Info("server exit properly")
}
