package options

import (
	"github.com/spf13/cobra"
)

type WorkerRunOptions struct {
	Port int
}

func NewWorkerRunOptions(cmd *cobra.Command) *WorkerRunOptions {
	opts := &WorkerRunOptions{}
	cmd.Flags().IntVarP(&opts.Port, "port", "p", 8080, "")
	return opts
}
