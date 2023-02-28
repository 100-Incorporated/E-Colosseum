package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var PATH string = "./users.db"

type User struct {
	Id       int
	Username string
	Password     string
}

func create_table() {
	//open database
	db, err := sql.Open("sqlite3", PATH)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//create table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, username TEXT, password TEXT)")
	if err != nil {
		panic(err)
	}

}

func clear_table() {
	//open database
	db, err := sql.Open("sqlite3", PATH)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//clear table
	_, err = db.Exec("DELETE FROM users")
	if err != nil {
		panic(err)
	}
}


func add_user(username string, password string) {
	//open database
	db, err := sql.Open("sqlite3", PATH)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//insert user
	_, err = db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
	if err != nil {
		panic(err)
	}
}


func show_users(path string) {
	//open database
	db, err := sql.Open("sqlite3", PATH)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//select all users
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	
	//for each row, print user info
	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Username, &user.Password)
		fmt.Println(user.Id, user.Username, user.Password)
	}
}


func main() {
	clear_table()
	create_table()
	add_user("root", "root")
	add_user("admin", "admin")
	show_users(PATH)
}


