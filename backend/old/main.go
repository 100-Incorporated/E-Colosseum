package main

import (
	"log"
	"net/http"
)

func main() {

	// This is to test the backend GET function, open http://localhost:80/users/1 in your browser
	users = append(users, User{
		ID:       1,
		Username: "JohnDoe",
		Rank:     1,
	})

	host := "127.0.0.1:80"
	if err := http.ListenAndServe(host, httpHandler()); err != nil {
		log.Fatalf("Failed to listen on %s: %v", host, err)
	}

}
