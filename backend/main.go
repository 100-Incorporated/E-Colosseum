package main

import (
	"log"
	"net/http"
)

func main() {

	users = append(users, User{
		ID:       1,
		Username: "JohnDoe",
		Rank:     1,
	})

	host := "127.0.0.1:82"
	if err := http.ListenAndServe(host, httpHandler()); err != nil {
		log.Fatalf("Failed to listen on %s: %v", host, err)
	}

}
