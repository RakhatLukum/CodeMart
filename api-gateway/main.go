package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	// TODO: Register REST endpoints that proxy to gRPC services
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("api-gateway is running"))
	})

	log.Println("API Gateway running on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("failed to start api-gateway: %v", err)
	}
}
