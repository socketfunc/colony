package options

import (
	"github.com/spf13/cobra"
)

type WorkerRunOptions struct {
	Port      int
	ProxyHost string
}

func NewWorkerRunOptions(cmd *cobra.Command) *WorkerRunOptions {
	opts := &WorkerRunOptions{}
	cmd.Flags().IntVarP(&opts.Port, "port", "p", 8080, "")
	cmd.Flags().StringVarP(&opts.ProxyHost, "proxy-host", "h", "localhost", "host of grpc proxy server")
	return opts
}
