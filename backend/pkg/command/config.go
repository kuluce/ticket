package command

import (
	"fmt"
	"oneTicket/backend/internal/config"

	"github.com/spf13/cobra"
)

var (
	configCmd = &cobra.Command{
		Use:   "config",
		Short: "config",
		Long:  `config.`,
		Run: func(cmd *cobra.Command, args []string) {
			configurate()
		},
	}
	configAction string
)

func configurate() {
	fmt.Printf("configAction: %v\n", configAction)

	switch configAction {
	case "new":
		newConfig()
	case "read":
		readconfig()
	case "save":
		saveConfig()
	default:
		readconfig()
	}

}
func init() {
	// configCmd.Flags().StringVarP(configAction, "new", "new", "new", "new")
	configCmd.Flags().StringVarP(&configAction, "action", "a", "", "Action to perform")
}

// 新建配置
func newConfig() {
	if err := config.New(); err != nil {
		fmt.Printf("新建配置失败，err: %v\n", err)
	}
	saveConfig()
}

// 读取配置
func readconfig() {
	if err := config.Read(); err != nil {
		fmt.Printf("获取配置失败，err: %v\n", err)
	}

	if err := config.Show(); err != nil {
		fmt.Printf("显示配置失败，err: %v\n", err)
	}

}

func saveConfig() {
	if err := config.Save(); err != nil {
		fmt.Printf("保存配置失败，err: %v\n", err)
	}
}
