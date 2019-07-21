package app

import (
	"fmt"
	"net"

	"github.com/socketfunc/colony/cmd/colony-worker/app/options"
	"github.com/socketfunc/colony/pkg/worker"
	"github.com/spf13/cobra"
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
	lis, err := net.Listen(network, fmt.Sprintf(":%d", opts.Port))
	if err != nil {
		return err
	}
	srv := grpc.NewServer()
	worker.RegisterWorkerServer(srv)
	return srv.Serve(lis)
}
