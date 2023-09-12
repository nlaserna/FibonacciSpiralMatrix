package main

import (
	"fibonacciSpiral/services"
	"log"
	"net/http"
)

func handleRequests() {
	http.HandleFunc("/", services.SpiralMatrixHandler)
	http.HandleFunc("/register", services.RegisterHandler)
	http.HandleFunc("/login", services.LoginHandler)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
