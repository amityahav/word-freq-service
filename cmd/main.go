package main

import (
	"context"
	"flag"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"wordStore/api"
	"wordStore/internal"
)

func main() {
	var cfgPath string

	flag.StringVar(&cfgPath, "config", "", "config path for service")
	flag.Parse()

	log.Println("startup: reading configuration file")
	cfg, err := readConfig(cfgPath)
	if err != nil {
		log.Fatalf("startup: %v", err)
	}

	err = validateConfig(cfg)
	if err != nil {
		log.Fatalf("startup: %v", err)
	}

	a := api.NewAPI(cfg)

	defer a.Shutdown(context.Background())

	log.Printf("serving at address: %s", cfg.ListenAddress)
	if err = a.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func readConfig(path string) (*internal.Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg internal.Config

	err = yaml.Unmarshal(b, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func validateConfig(cfg *internal.Config) error {
	if cfg.Store.K <= 0 {
		return ErrInvalidK
	}

	return nil
}

var ErrInvalidK = errors.New("K must be greater than zero")
