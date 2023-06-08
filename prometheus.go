package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	motdRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "motd_requests_received",
		Help: "The total number of requests",
	})

	motdServed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "motd_messages_served",
		Help: "The total number of messages served",
	})
)

func StartMetrics() {
	go func() {
		port := fmt.Sprintf(":%d", 2112)

		log.Printf("Starting metrics server on %s", port)

		s := &http.Server{
			Addr:           port,
			Handler:        promhttp.Handler(),
			ReadTimeout:    60 * time.Second,
			WriteTimeout:   60 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}

		log.Fatal(s.ListenAndServe())
	}()
}
