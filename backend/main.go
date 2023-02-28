package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Rank     int    `json:"rank"`
}

// Users slice to test functionality
var users = []User{
	{ID: "1", Username: "JohnDoe", Rank: 1},
	{ID: "2", Username: "JaneDoe", Rank: 2},
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)

	router.Run(":8080")
}

// getUsers returns a slice of all users as JSON
func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}
