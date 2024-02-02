package util

import (
	"os"
	"os/signal"
	"syscall"
)

func ReadPortArg() string {
	var port = ":8080"
	if len(os.Args) < 2 {
		return port
	}
	rawPort := os.Args[1]
	if rawPort != "" {
		if rawPort[0] != ':' {
			port = ":" + rawPort
		}
	}
	return port
}

func NewStopListener() chan os.Signal {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT)
	return signalChan
}

func WebAssetPath() string {
	return `..\dist\src`
}
