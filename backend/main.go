package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Rank     int    `json:"rank"`
}

func main() {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

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
	router.GET("/users", func(c *gin.Context) {
		rows, err := db.Query("SELECT * FROM users")
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}
		defer rows.Close()

		users := []User{}
		for rows.Next() {
			var user User
			err := rows.Scan(&user.ID, &user.Username, &user.Rank)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
				return
			}
			users = append(users, user)
		}

		c.IndentedJSON(http.StatusOK, users)
	})

	router.GET("/users/:id", func(c *gin.Context) {
		var user User
		err := db.QueryRow("SELECT * FROM users WHERE id=?", c.Param("id")).Scan(&user.ID, &user.Username, &user.Rank)
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
			return
		} else if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}

		c.IndentedJSON(http.StatusOK, user)
	})

	router.POST("/users", func(c *gin.Context) {
		var newUser User
		if err := c.BindJSON(&newUser); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
			return
		}

		_, err := db.Exec("INSERT INTO users (username, rank) VALUES (?, ?)", newUser.Username, newUser.Rank)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}

		c.IndentedJSON(http.StatusCreated, newUser)
	})

	router.Run("localhost:8081")
}
