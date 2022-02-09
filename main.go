package main

import (
	"github.com/pwood/middleauth/check"
	"github.com/pwood/middleauth/check/iplist"
	"github.com/pwood/middleauth/check/static"
	"github.com/pwood/middleauth/http"
	"log"
	"os"
	"os/signal"
)

func main() {
	ipList := []string{"10.0.0.0/8", "172.16.0.0/12", "192.168.0.0/16"}
	i, err := iplist.New(ipList, check.ACCEPT)
	if err != nil {
		log.Panicf("failed to create IP list: %s", err.Error())
	}

	s, err := static.New(check.REJECT)
	if err != nil {
		log.Panicf("failed to create static: %s", err.Error())
	}

	checks := []check.Checker{i, s}

	srv := http.Server{
		Port:   8080,
		Checks: checks,
	}

	serverDone, err := srv.Start()
	if err != nil {
		log.Panicf("failed to start http server: %s", err.Error())
	}

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, os.Kill)

	sig := <-signalCh

	log.Printf("signal received, shutting down: %d", sig)

	if err := serverDone(); err != nil {
		log.Panicf("failed to stop http server: %s", err.Error())
	}
}
