package main

import (
	"github.com/nasjp/golanglab/queueworker/queueworker"
	"github.com/nasjp/golanglab/queueworker/services"
)

const (
	dispatcherN = 3
	workerN     = 10
)

func main() {
	qw := queueworker.New(dispatcherN, workerN, services.SampleHandler)
	qw.Run()
}
