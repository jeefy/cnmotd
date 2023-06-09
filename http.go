package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func StartHTTP() {
	port := fmt.Sprintf(":%d", args.httpPort)
	log.Printf("Starting HTTP server on %s", port)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	http.Handle("/", http.HandlerFunc(handleRequest))
	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", args.httpPort), nil); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Print("Server Started")

	<-done
	log.Print("Server Stopped")
}
