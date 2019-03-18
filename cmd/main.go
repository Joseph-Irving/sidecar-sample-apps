package main

import (
	"os"
)

func main() {
	arg := os.Args[1]

	switch arg {
	case "sidecar":
		runSidecar()
	case "worker":
		runWorker()
	case "primary":
		runPrimary()
	}
}
