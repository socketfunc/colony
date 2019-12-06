package app

import (
	"context"
	"fmt"
	"net"

	"github.com/socketfunc/colony/cmd/colony-worker/app/options"
	"github.com/socketfunc/colony/pkg/auth"
	"github.com/socketfunc/colony/pkg/log"
	"github.com/socketfunc/colony/pkg/storage"
	"github.com/socketfunc/colony/pkg/worker"
	apiv1beta1 "github.com/socketfunc/colony/proto/api/v1beta1"
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger, err := log.NewLogger()
	if err != nil {
		return err
	}
	defer logger.Sync()
	lis, err := net.Listen(network, fmt.Sprintf(":%d", opts.Port))
	if err != nil {
		logger.Error("Failed to listen port",
			zap.Error(err), zap.Int("port", opts.Port))
		return err
	}

	conn, err := createGRPCClient(ctx, opts.ProxyHost)
	if err != nil {
		logger.Error("Failed to create grpc client",
			zap.Error(err), zap.String("proxy-host", opts.ProxyHost))
		return err
	}

	params := &worker.ServerParams{
		ApiClient: apiv1beta1.NewApiServiceClient(conn),
		Storage:   storage.New(),
	}
	srv := grpc.NewServer(
		grpc.UnaryInterceptor(
			auth.AuthorizeInterceptor(),
		),
	)
	worker.RegisterWorkerServer(
		srv,
		params,
		worker.WithLogger(logger),
	)
	logger.Info("Start worker")
	return srv.Serve(lis)
}

func createGRPCClient(ctx context.Context, host string) (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{
		grpc.WithUserAgent("colony.worker.v1beta1"),
	}
	return grpc.DialContext(ctx, host, opts...)
}
