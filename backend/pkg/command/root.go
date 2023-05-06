package command

import (
	"oneTicket/backend/internal/server"

	"github.com/spf13/cobra"
)

// 子命令注册
func init() {
	RootCmd.AddCommand(RunCmd)
	RootCmd.AddCommand(InstallCmd)
	RootCmd.AddCommand(unInstallCmd)
	RootCmd.AddCommand(StartCmd)
	RootCmd.AddCommand(StopCmd)
	RootCmd.AddCommand(configCmd)
}

var (
	RootCmd = &cobra.Command{
		Use:   "oneTicket",
		Short: "oneTicket",
		RunE: func(cmd *cobra.Command, args []string) error {
			server.Start()
			return nil
		},
	}
)
