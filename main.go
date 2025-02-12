package main

import (
	_ "embed"
	"os"

	"github.com/spf13/cobra"
	"github.com/weedbox/common-modules/configs"

	"go.uber.org/fx"
)

const (
	appName        = "template-project"
	appDescription = "template-project is a general service."
)

// Global variables
var config *configs.Config

// Command options
var pc bool
var verbose bool

func main() {

	rootCmd := &cobra.Command{
		Use:  appName,
		Long: appDescription,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := run(); err != nil {
				return err
			}
			return nil
		},
	}

	// Initiaizing config and Command line flags
	config = configs.NewConfig("SERVICE")
	rootCmd.Flags().BoolVar(&pc, "print_configs", false, "Print all available configs")
	rootCmd.Flags().BoolVar(&verbose, "verbose", false, "Display detailed logs")

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initModules() ([]fx.Option, error) {

	modules := []fx.Option{}

	// preload
	m, err := preloadModules()
	if err != nil {
		return modules, err
	}

	modules = append(modules, m...)

	// general
	m, err = loadModules()
	if err != nil {
		return modules, err
	}

	modules = append(modules, m...)

	// After
	m, err = afterModules()
	if err != nil {
		return modules, err
	}

	modules = append(modules, m...)

	if !verbose {
		modules = append(modules, fx.NopLogger)
	}

	return modules, nil
}

func run() error {

	modules, err := initModules()
	if err != nil {
		return err
	}

	app := fx.New(
		modules...,
	)

	if pc {
		config.PrintAllSettings()
	}

	app.Run()

	return nil
}
