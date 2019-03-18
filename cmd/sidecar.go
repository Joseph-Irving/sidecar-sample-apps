package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func runSidecar() {

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)
	signal.Notify(stopChan, syscall.SIGTERM)

	wakeup := 0

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		if wakeup >= 5 {
			w.WriteHeader(200)
		} else {
			w.WriteHeader(503)
		}
	})
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to listen and serve server: %v", err)
		}
	}()

	ticker := time.Tick(2 * time.Second)

	for {
		select {
		case <-stopChan:
			fmt.Println("I'M DYING")

			ctx, cancel := context.WithTimeout(context.Background(), (5 * time.Second))
			defer cancel()

			server.Shutdown(ctx)

			time.Sleep(2 * time.Second)
			fmt.Println("I'M DEAD")
			os.Exit(0)
		case <-ticker:
			if wakeup < 5 {
				fmt.Println("I'M NOT READY!")
				wakeup++
			} else {
				fmt.Println("I'M READY!")
			}
		}
	}

}
