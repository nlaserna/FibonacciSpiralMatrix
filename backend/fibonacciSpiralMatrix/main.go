package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

func homePage(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	switch r.Method {
	case "OPTIONS":
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		return
	}
	rows, _ := strconv.Atoi(r.URL.Query().Get("rows"))
	columns, _ := strconv.Atoi(r.URL.Query().Get("columns"))

	v := SpiralMatrix(rows, columns)
	matrixResponse := new(MatrixResponse)
	matrixResponse.Timestamp = time.Now().Unix()
	matrixResponse.Rows = v
	json.NewEncoder(w).Encode(matrixResponse)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/register", RegisterHandler)
	http.HandleFunc("/login", LoginHandler)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
