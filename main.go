package main

import (
	"fmt"
	"log"
	"net/http"
)

// helloHandler handles requests to the "/" route
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the path is exactly "/" to avoid matching everything
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World! Welcome to your Go web server v3.")
}

// apiHandler handles requests to the "/api" route
func apiHandler(w http.ResponseWriter, r *http.Request) {
	// Set the response header to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Success", "status": "active"}`))
}

func main() {
	// 1. Create a new request multiplexer (router)
	mux := http.NewServeMux()

	// 2. Register handler functions for specific routes
	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/api", apiHandler)

	// 3. Define the server address
	port := ":8080"
	fmt.Printf("Starting server on http://localhost%s...\n", port)

	// 4. Start the server and log if it crashes
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
