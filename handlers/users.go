package handlers

import (
	"encoding/json"
	"net/http"
)

// User struct with correct JSON field tags
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// UserHandler function to handle HTTP requests
func UserHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name:  "John Doe",
		Email: "john.doe@example.com",
		Age:   30,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
