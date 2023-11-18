package main

import (
	"log/slog"

	"github.com/henderiw-nephio/kform/kform-plugin/kfprotov1"
	"github.com/henderiw-nephio/kform/kform-plugin/kfprotov1/kfserver1"
	"github.com/henderiw-nephio/kform/kform-sdk-go/pkg/schema"
	"github.com/henderiw/logger/log"
	resourcebackend "github.com/kformdev/provider-resourcebackend/provider"
)

const providerName = "registry.kform.dev/provider/resourcebackend"

func main() {

	log := log.NewLogger(&log.HandlerOptions{Name: "provider-resourcebackend-logger", AddSource: false})
	slog.SetDefault(log)

	grpcProviderFunc := func() kfprotov1.ProviderServer {
		return schema.NewGRPCProviderServer(resourcebackend.Provider())
	}

	opts := []kfserver1.ServeOpt{
		kfserver1.WithGoPluginLogger(log),
	}
	if err := kfserver1.Serve(
		providerName,
		grpcProviderFunc,
		opts...); err != nil {
		slog.Error("kform serve failed", "err", err)
	}
	log.Info("done serving kform")

}
