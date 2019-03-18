package main

import (
	"fmt"
	"os"
	"time"
)

func runWorker() {
	ticker := time.Tick(2 * time.Second)
	working := 0

	for {
		select {
		case <-ticker:
			if working < 5 {
				fmt.Println("I'M WORKING!")
				working++
			} else {
				fmt.Println("I'M DONE!")
				os.Exit(0)
			}
		}
	}

}
