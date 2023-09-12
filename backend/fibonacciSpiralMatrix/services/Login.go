package services

import (
	"encoding/json"
	"fibonacciSpiralMatrix/structs"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var users = make(map[string]structs.User)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	var creds structs.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, exists := users[creds.Username]
	if !exists || !checkPasswordHash(creds.Password, user.Password) {
		http.Error(w, "Authentication failed", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Login successful"}
	json.NewEncoder(w).Encode(response)
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
