package command

import (
	"ticket/backend/internal/server"

	"github.com/spf13/cobra"
)

var (
	RunCmd = &cobra.Command{
		Use:   "run",
		Short: "run service",
		Long:  `run service`,
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func run() {
	server.Start()
}
