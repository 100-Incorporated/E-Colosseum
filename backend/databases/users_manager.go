package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id       int
	Username string
	Password     string
}

func main() {
	//open database
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//create table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, username TEXT, password TEXT)")
	if err != nil {
		panic(err)
	}
	
	//insert a user username Jeff password 123
	db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", "Jeff", "123")
	

	//insert a user username Brayden password 456
	db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", "Brayden", "456")
	
	//select all users
	rows, err := db.Query("SELECT * FROM users")
	
	//for each row, print user info
	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Username, &user.Password)
		fmt.Println(user.Id, user.Username, user.Password)
	}

}


