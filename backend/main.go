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
	router.GET("/users/:id", getUserByID)
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

// getUserByID returns a single user based on their id
func getUserByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of users, looking for
	// a user whose ID value matches the parameter.
	for _, a := range users {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}
