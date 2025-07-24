package main

import (
	"log"
	"net/http"

	"{{.AppName}}/internal/router"
	"{{.AppName}}/internal/middleware"
)

func main() {
	mux := http.NewServeMux()

	router.RegisterRoutes(mux)

	handler := middleware.ApplyMiddlewares(mux)

	port := ":8080"
	log.Printf("Starting server on %s...\n", port)
	if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
