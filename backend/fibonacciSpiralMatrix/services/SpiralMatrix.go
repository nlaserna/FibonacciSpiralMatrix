package services

import (
	"encoding/json"
	"fibonacciSpiralMatrix/helper"
	"fibonacciSpiralMatrix/structs"
	"net/http"
	"strconv"
	"time"
)

func SpiralMatrixHandler(w http.ResponseWriter, r *http.Request) {
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

	v := helper.SpiralMatrix(rows, columns)
	matrixResponse := new(structs.MatrixResponse)
	matrixResponse.Timestamp = time.Now().Unix()
	matrixResponse.Rows = v
	json.NewEncoder(w).Encode(matrixResponse)
}
