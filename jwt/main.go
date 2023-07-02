package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key")

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

func login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")

	// Authenticate the user, for example, by checking credentials against a database
	// ...

	// Create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Fatal(err)
	}

	// Return the token as a response
	fmt.Fprintf(w, tokenString)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/login", login)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
