package main

import (
	"context"
	"github.com/pwood/middleauth/check"
	"github.com/pwood/middleauth/check/iplist"
	"github.com/pwood/middleauth/check/static"
	"github.com/pwood/middleauth/http"
	"github.com/sethvargo/go-envconfig"
	"log"
	"os"
	"os/signal"
)

type Config struct {
	ConfigMode string `env:"CONFIG_MODE,default=ENV"`
	ConfigFile string `env:"CONFIG_FILE"`

	ServerHost string `env:"SERVER_HOST"`
	ServerPort int    `env:"SERVER_PORT,default=8888"`

	PermittedNetworks []string `env:"PERMITTED_NETWORKS"`
	PermittedMTLSAny  bool     `env:"PERMITTED_MTLS_ANY,default=false"`
}

func main() {
	ctx := context.Background()

	cfg := Config{}
	if err := envconfig.Process(ctx, &cfg); err != nil {
		log.Panicf("failed to parse environment for config: %s", err.Error())
	}

	var srv http.Server

	if cfg.ConfigMode == "ENV" {
		srv = constructServerFromEnvironmentConfig(cfg)
	} else {
		log.Panicf("unacceptable config mode: %s", cfg.ConfigMode)
	}

	serverDone, err := srv.Start()
	if err != nil {
		log.Panicf("failed to start http server: %s", err.Error())
	}

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, os.Kill)

	sig := <-signalCh

	log.Printf("signal received, shutting down: %d", sig)

	if err := serverDone(ctx); err != nil {
		log.Panicf("failed to stop http server: %s", err.Error())
	}
}

func constructServerFromEnvironmentConfig(cfg Config) http.Server {
	checks := []check.Checker{}

	i, err := iplist.New(cfg.PermittedNetworks, check.ACCEPT)
	if err != nil {
		log.Panicf("failed to create IP list: %s", err.Error())
	}

	checks = append(checks, i)

	if cfg.PermittedMTLSAny {
		m, err := static.New(check.ACCEPT)
		if err != nil {
			log.Panicf("failed to create mtls: %s", err.Error())
		}

		checks = append(checks, m)
	}

	s, err := static.New(check.REJECT)
	if err != nil {
		log.Panicf("failed to create static: %s", err.Error())
	}

	checks = append(checks, s)

	srv := http.Server{
		Port:   cfg.ServerPort,
		Host:   cfg.ServerHost,
		Checks: checks,
	}

	return srv
}
