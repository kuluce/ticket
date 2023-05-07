package main

import (
	"fmt"
	"os"

	"ticket/backend/pkg/command"

	"github.com/toolkits/pkg/logger"
)

func main() {
	logger.Info("ticket")

	if err := command.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
