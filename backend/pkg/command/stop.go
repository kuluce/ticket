package command

import (
	"fmt"
	"oneTicket/backend/pkg/srv"

	"github.com/spf13/cobra"
)

var (
	StopCmd = &cobra.Command{
		Use:   "stop",
		Short: "stop service",
		Long:  `stop service`,
		Run: func(cmd *cobra.Command, args []string) {
			stop()
		},
	}
)

func stop() {
	if err := srv.Stop(); err != nil {
		fmt.Printf("stop service failed,err: %v\n", err)
	}
}
