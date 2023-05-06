package srv

import (
	"fmt"

	"github.com/kardianos/service"
	"github.com/toolkits/file"
	"github.com/toolkits/logger"
)

type program struct{}

// Start supervised service
func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {}

// Stop supervised service
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

func serviceOperation(action string) error {

	installedPath := file.SelfDir()
	serviceArgs := make([]string, 0)
	serviceArgs = append(serviceArgs, "-f="+fmt.Sprintf("%s/one.yml", installedPath))

	svcConfig := &service.Config{
		Name:             "ticketd",
		DisplayName:      "ticketd",
		Description:      "ticket service",
		WorkingDirectory: installedPath,
		UserName:         "root",
		Arguments:        serviceArgs,
		Dependencies: []string{
			"After=network-online.target",
			"Wants=network-online.target",
		},
	}
	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		logger.Error("service init failed", err)
		return err
	}
	switch action {
	case "install":
		err := s.Install()
		if err != nil {
			logger.Error("Failed to install service ticked: ", err)
			fmt.Println("Failed to install service ticked: ", err)
			return err
		} else {
			fmt.Println("Succeed to install service ticked")
		}
	case "uninstall":
		s.Stop()
		err := s.Uninstall()
		if err != nil {
			logger.Error("Failed to uninstall service ticked: ", err)
			fmt.Println("Failed to uninstall service ticked: ", err)
			return err
		} else {
			fmt.Println("Succeed to uninstall service ticked")
		}
	case "start":
		err := s.Start()
		if err != nil {
			logger.Error("Failed to start service: ", err)
			fmt.Println("Failed to start service: ", err)
			return err
		} else {
			fmt.Println("Succeed to start service ticked")
		}
	case "stop":
		err := s.Stop()
		if err != nil {
			logger.Error("Failed to stop service: ", err)
			fmt.Println("Failed to stop service: ", err)
			return err
		} else {
			fmt.Println("Succeed to stop service ticked")
		}

	default:
		fmt.Println("install/uninstall/start/stop service")
	}
	return nil
}
