package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Atsu-Imo/goroutine-channel-learning/server"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	// Ignore all signals
	signal.Ignore()
	signal.Notify(sigChan, syscall.SIGINT)

	listener := server.NewListener()
	go listener.Listen()
	fmt.Println("listener started")
	select {
	case sig := <-sigChan:
		switch sig {
		case syscall.SIGINT:
			fmt.Println("shutdown")
		default:
			fmt.Println("unexpected signal")
		}
	}
}
