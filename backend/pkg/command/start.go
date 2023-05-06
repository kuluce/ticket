package command

import (
	"fmt"
	"oneTicket/backend/pkg/srv"

	"github.com/spf13/cobra"
)

var (
	StartCmd = &cobra.Command{
		Use:   "start",
		Short: "start service",
		Long:  `start service`,
		Run: func(cmd *cobra.Command, args []string) {
			start()
		},
	}
)

func start() {
	if err := srv.Start(); err != nil {
		fmt.Printf("start service failed,err: %v\n", err)
	}
}
