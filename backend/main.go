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
	Password string `json:"password"`
	Birthday string `json:"birthday"`
}

func main() {
	// Stop the program if the database connection fails
	db, err := sql.Open("sqlite3", "./databases/users.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Using default middleware (logger and recovery)
	router := gin.Default()

	// Enable on release?
	// gin.SetMode(gin.ReleaseMode)

	//CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"X-Requested-With", "Content-Type", "Authorization",
		"DNT", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since",
		"Cache-Control", "Content-Range", "Range"}
	config.ExposeHeaders = []string{"DNT", "Keep-Alive", "User-Agent",
		"X-Requested-With", "If-Modified-Since", "Cache-Control",
		"Content-Type", "Content-Range", "Range", "Content-Disposition"}
	config.AllowCredentials = true
	config.MaxAge = 86400

	// router.Use(cors.New(config))

	// Routes
	// GET all users
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
			err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Birthday)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
				return
			}
			users = append(users, user)
		}

		c.IndentedJSON(http.StatusOK, users)
	})

	// GET a user by ID
	router.GET("/users/:id", func(c *gin.Context) {
		var user User
		err := db.QueryRow("SELECT * FROM users WHERE id=?", c.Param("id")).Scan(&user.ID, &user.Username, &user.Password, &user.Birthday)
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		} else if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}

		c.IndentedJSON(http.StatusOK, user)
	})

	// POST a new user
	router.POST("/users", func(c *gin.Context) {
		var newUser User
		if err := c.BindJSON(&newUser); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
			return
		}

		// note that the value of the id field is automatically created by the database
		result, err := db.Exec("INSERT INTO users (username, password, birthday) VALUES (?, ?, ?)", newUser.Username, newUser.Password, newUser.Birthday)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}

		// set newUser id to the id of the newly created user (otherwise it will be 0)
		var id, _ = result.LastInsertId()
		newUser.ID = int(id)

		c.IndentedJSON(http.StatusCreated, newUser)
	})

	// DELETE a user by ID
	router.DELETE("/users/:id", func(c *gin.Context) {
		userID := c.Param("id")
		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM users WHERE id = ?", userID).Scan(&count)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}

		if count == 0 {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		}

		_, err = db.Exec("DELETE FROM users WHERE id = ?", userID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	})

	// PUT a user by ID
	router.PUT("/users/:id", func(c *gin.Context) {
		// The user variable just exists to store the result of the query, so the scan output can be stored in err
		var user User
		err := db.QueryRow("SELECT * FROM users WHERE id=?", c.Param("id")).Scan(&user.ID, &user.Username, &user.Password, &user.Birthday)
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		} else if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}

		var updatedUser User
		if err := c.BindJSON(&updatedUser); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
			return
		}

		result, err := db.Exec("UPDATE users SET username=?, password=?, birthday=? WHERE id=?", updatedUser.Username, updatedUser.Password, updatedUser.Birthday, c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}

		// set updatedUser id to the id of the newly created user
		var id, _ = result.LastInsertId()
		updatedUser.ID = int(id)

		c.IndentedJSON(http.StatusOK, updatedUser)
	})

	router.PATCH("/users/:id", func(c *gin.Context) {
		var user User
		err := db.QueryRow("SELECT * FROM users WHERE id=?", c.Param("id")).Scan(&user.ID, &user.Username, &user.Password, &user.Birthday)
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
			return
		} else if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}

		var updatedUser User
		if err := c.BindJSON(&updatedUser); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
			return
		}
		// Set fields that are not provided to their current values
		if updatedUser.Username == "" {
			updatedUser.Username = user.Username
		}
		if updatedUser.Password == "" {
			updatedUser.Password = user.Password
		}
		if updatedUser.Birthday == "" {
			updatedUser.Birthday = user.Birthday
		}

		result, err := db.Exec("UPDATE users SET username=?, password=?, birthday=? WHERE id=?", updatedUser.Username, updatedUser.Password, updatedUser.Birthday, c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}

		// set updatedUser id to the id of the newly created user
		var id, _ = result.LastInsertId()
		updatedUser.ID = int(id)

		c.IndentedJSON(http.StatusOK, updatedUser)
	})

	router.Run("localhost:8080")
}
