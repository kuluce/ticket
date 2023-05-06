package command

import (
	"fmt"
	"oneTicket/backend/pkg/srv"

	"github.com/spf13/cobra"
)

var (
	InstallCmd = &cobra.Command{
		Use:   "install",
		Short: "install service",
		Long:  `install service`,
		Run: func(cmd *cobra.Command, args []string) {
			install()
		},
	}
)

func install() {
	if err := srv.Install(); err != nil {
		fmt.Printf("install service failed,err: %v\n", err)
	}
}
