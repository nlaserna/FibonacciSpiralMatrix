package main

import (
	"fibonacciSpiralMatrix/services"
	"github.com/rs/cors"
	"net/http"
)

func handleRequests() {
	mux := http.NewServeMux()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowOriginFunc: func(origin string) bool { return true },
		Debug: true,
	})
	handler := c.Handler(mux)
	mux.HandleFunc("/", services.SpiralMatrixHandler)
	mux.HandleFunc("/register", services.RegisterHandler)
	mux.HandleFunc("/login", services.LoginHandler)
	http.ListenAndServe(":10000", handler)
}

func main() {
	handleRequests()
}
