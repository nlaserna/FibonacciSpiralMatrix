package services

import (
	"encoding/json"
	"fibonacciSpiralMatrix/structs"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var users = make(map[string]structs.User)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user structs.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var foundUser structs.User
	for _, u := range registeredUsers {
		if u.Username == user.Username {
			foundUser = u
			break
		}
	}

	if foundUser.Username == "" {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Login successful")
}
