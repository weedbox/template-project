package main

import (
	"github.com/weedbox/common-modules/daemon"
	"github.com/weedbox/common-modules/logger"
	"go.uber.org/fx"
)

func preloadModules() ([]fx.Option, error) {
	modules := []fx.Option{
		fx.Supply(config),
		logger.Module(),
	}

	return modules, nil
}

func afterModules() ([]fx.Option, error) {
	modules := []fx.Option{
		daemon.Module("daemon"),
	}

	return modules, nil
}

func loadModules() ([]fx.Option, error) {
	modules := []fx.Option{}

	return modules, nil
}
