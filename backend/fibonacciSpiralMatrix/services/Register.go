package services

import (
	"encoding/json"
	"fibonacciSpiralMatrix/structs"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

var registeredUsers []structs.User

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user structs.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, u := range registeredUsers {
		if u.Username == user.Username {
			http.Error(w, "Username already exists", http.StatusConflict)
			return
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	user.Password = string(hashedPassword)
	registeredUsers = append(registeredUsers, user)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Registration successful")
}
