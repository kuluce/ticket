package command

import (
	"fmt"
	"oneTicket/backend/pkg/srv"

	"github.com/spf13/cobra"
)

var (
	unInstallCmd = &cobra.Command{
		Use:   "uninstall",
		Short: "uninstall",
		Long:  `uninstall.`,
		Run: func(cmd *cobra.Command, args []string) {
			uninstall()
		},
	}
)

func uninstall() {
	if err := srv.Uninstall(); err != nil {
		fmt.Printf("uninstall service failed,err: %v\n", err)
	}
}
