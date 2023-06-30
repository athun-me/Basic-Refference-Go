package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	// Define routes and their handlers
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/user", getUser)

	// Start the server
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}

func getUser(w http.ResponseWriter, r *http.Request) {
	user := struct {
		User string `json:"user"`
	}{
		User: "This is JSON response",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
