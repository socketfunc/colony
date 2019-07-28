package app

import (
	"fmt"
	"net"

	"github.com/socketfunc/colony/cmd/colony-worker/app/options"
	"github.com/socketfunc/colony/pkg/log"
	"github.com/socketfunc/colony/pkg/worker"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	network = "tcp"
	cliName = "colony-worker"
)

func NewAppCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: cliName,
	}
	opts := options.NewWorkerRunOptions(cmd)
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return run(opts)
	}
	return cmd
}

func run(opts *options.WorkerRunOptions) error {
	logger, err := log.NewLogger()
	if err != nil {
		return err
	}
	defer logger.Sync()
	lis, err := net.Listen(network, fmt.Sprintf(":%d", opts.Port))
	if err != nil {
		logger.Error("Failed to listen port", zap.Error(err), zap.Int("port", opts.Port))
		return err
	}
	params := &worker.ServerParams{}
	srv := grpc.NewServer()
	worker.RegisterWorkerServer(
		srv,
		params,
		worker.WithLogger(logger),
	)
	logger.Info("Start worker")
	return srv.Serve(lis)
}
