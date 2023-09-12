package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

// Credentials represent the login request payload.
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = make(map[string]User)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Check if the username is already taken
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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
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

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
