package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Birthday string `json:"birthday" binding:"required"`
}

type UserResponse struct {
	ID       int    `json:"id" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Birthday string `json:"birthday" binding:"required"`
}

type UserUpdate struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Birthday string `json:"birthday"`
}

func setupRouter(db *sql.DB) *gin.Engine {
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

	router.Use(cors.New(config))

	// Routes
	// GET all users
	router.GET("/users", func(c *gin.Context) {
		rows, err := db.Query("SELECT * FROM users")
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}
		defer rows.Close()

		users := []UserResponse{}
		for rows.Next() {
			var user UserResponse
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
		var user UserResponse
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
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
			return
		}

		// note that the value of the id field is automatically created by the database
		result, err := db.Exec("INSERT INTO users (username, password, birthday) VALUES (?, ?, ?)", newUser.Username, newUser.Password, newUser.Birthday)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}

		var id, _ = result.LastInsertId()
		createdUser := UserResponse{
			ID:       int(id),
			Username: newUser.Username,
			Password: newUser.Password,
			Birthday: newUser.Birthday,
		}

		c.IndentedJSON(http.StatusCreated, createdUser)
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
		var user UserResponse
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
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
			return
		}

		result, err := db.Exec("UPDATE users SET username=?, password=?, birthday=? WHERE id=?", updatedUser.Username, updatedUser.Password, updatedUser.Birthday, c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}

		// set updatedUser id to the id of the newly created user
		var id, _ = result.LastInsertId()
		createdUser := UserResponse{
			ID:       int(id),
			Username: updatedUser.Username,
			Password: updatedUser.Password,
			Birthday: updatedUser.Birthday,
		}

		c.IndentedJSON(http.StatusOK, createdUser)
	})

	router.PATCH("/users/:id", func(c *gin.Context) {
		var user UserResponse
		err := db.QueryRow("SELECT * FROM users WHERE id=?", c.Param("id")).Scan(&user.ID, &user.Username, &user.Password, &user.Birthday)
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		} else if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}

		var updateFields UserUpdate
		if err := c.BindJSON(&updateFields); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
			return
		}
		// Set fields that are not provided to their current values
		if updateFields.Username == "" {
			updateFields.Username = user.Username
		}
		if updateFields.Password == "" {
			updateFields.Password = user.Password
		}
		if updateFields.Birthday == "" {
			updateFields.Birthday = user.Birthday
		}

		result, err := db.Exec("UPDATE users SET username=?, password=?, birthday=? WHERE id=?", updateFields.Username, updateFields.Password, updateFields.Birthday, c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}

		var id, _ = result.LastInsertId()
		createdUser := UserResponse{
			ID:       int(id),
			Username: updateFields.Username,
			Password: updateFields.Password,
			Birthday: updateFields.Birthday,
		}

		c.IndentedJSON(http.StatusOK, createdUser)
	})
	return router
}

func main() {
	// Stop the program if the database connection fails
	db, err := sql.Open("sqlite3", "./databases/users.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := setupRouter(db)

	router.Run("localhost:8080")
}
