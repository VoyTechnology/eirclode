package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"

	"eirclode.voy.technology/internal/cmd"
)

func main() {
	if err := run(); err != nil {
		log.Printf("%+v", err)
		os.Exit(1)
	}
}

func run() error {
	var cfg cmd.Config
	if err := envconfig.Process("eirclode", &cfg); err != nil {
		return fmt.Errorf("unable to parse config: %w", err)
	}

	return cmd.Run(cfg)
}
