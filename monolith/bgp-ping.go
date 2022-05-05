package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/sgrade/bgp-ping/monolith/ping"
)

var (
	// showVersion bool
	// version     string
	// gitCommit   string
	counter int // ping counter
	// timeout  string
	interval string // interval in time.Duration format, e.g. "1s"
	sigs     chan os.Signal
)

func main() {

	interval = "1s"
	counter = 4

	sigs = make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	var timeoutDuration time.Duration
	var intervalDuration time.Duration

	if res, err := strconv.Atoi(interval); err == nil {
		intervalDuration = time.Duration(res) * time.Millisecond
	} else {
		intervalDuration, err = time.ParseDuration(interval)
		if err != nil {
			fmt.Println("parse interval failed:\n", err)
			return
		}
	}

	var protocol ping.Protocol

	host := "10.0.12.11"
	port := 179

	parseHost, _ := ping.FormatIP(host)
	target := ping.Target{
		Timeout:  timeoutDuration,
		Interval: intervalDuration,
		Host:     parseHost,
		Port:     port,
		Counter:  counter,
		Protocol: protocol,
	}

	pinger := ping.NewTCPing()

	pinger.SetTarget(&target)
	pingerDone := pinger.Start()
	select {
	case <-pingerDone:
		break
	case <-sigs:
		break
	}

	fmt.Println(pinger.Result())
}
