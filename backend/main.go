package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
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

	// Enable on release?
	// gin.SetMode(gin.ReleaseMode)

	//CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8081"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"X-Requested-With", "Content-Type", "Authorization",
		"DNT", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since",
		"Cache-Control", "Content-Range", "Range"}
	config.ExposeHeaders = []string{"DNT", "Keep-Alive", "User-Agent",
		"X-Requested-With", "If-Modified-Since", "Cache-Control",
		"Content-Type", "Content-Range", "Range", "Content-Disposition"}
	config.AllowCredentials = true
	config.MaxAge = 86400

	router.Use(cors.New(config))

	// Routes
	router.GET("/users", getUsers)
	router.POST("/users", postUser)

	router.Run("localhost:8081")
}

// getUsers returns a slice of all users as JSON
func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

// postUser creates a new user from JSON received in the request body
func postUser(c *gin.Context) {
	var newUser User

	// Call BindJSON to bind the received JSON to
	// newUser.
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}
