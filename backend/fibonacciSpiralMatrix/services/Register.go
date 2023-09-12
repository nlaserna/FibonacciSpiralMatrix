package services

import (
	"encoding/json"
	"fibonacciSpiral/structs"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var newUser structs.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if _, exists := users[newUser.Username]; exists {
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}

	// Hash the user's password before storing it
	hashedPassword, err := hashPassword(newUser.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	newUser.Password = hashedPassword
	users[newUser.Username] = newUser

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Registration successful"}
	json.NewEncoder(w).Encode(response)
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
