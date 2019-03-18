package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func runPrimary() {

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)
	signal.Notify(stopChan, syscall.SIGTERM)
	ticker := time.Tick(2 * time.Second)

	for {
		select {
		case <-stopChan:
			fmt.Println("Shutting Down")
			time.Sleep(2 * time.Second)
			fmt.Println("Shut down complete")
			os.Exit(0)
		case <-ticker:
			fmt.Println("I'm doing stuff")
		}
	}

}
