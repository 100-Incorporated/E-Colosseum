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
	Birthday string
}

func create_table() {
	//open database
	db, err := sql.Open("sqlite3", PATH)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//create table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, username TEXT, password TEXT, birthday TEXT)")
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


func add_user(username string, password string, birthday string) string {
	//open database
	db, err := sql.Open("sqlite3", PATH)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//insert user
	_, err = db.Exec("INSERT INTO users (username, password, birthday) VALUES (?, ?, ?)", username, password, birthday)
	if err != nil {
		panic(err)
	} else {
		//return user info
		return username + " " + password + " " + birthday
	}
	
}

func get_user(id int) User {
	//open database
	db, err := sql.Open("sqlite3", PATH)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	//select user
	var user User
	err = db.QueryRow("SELECT * FROM users WHERE id=?", id).Scan(&user.Id, &user.Username, &user.Password, &user.Birthday)
	if err == sql.ErrNoRows {
		panic(err)
	} else if err != nil {
		panic(err)
	}

	return user

}

func delete_user(id int) {
	//open database
	db, err := sql.Open("sqlite3", PATH)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//delete user
	_, err = db.Exec("DELETE FROM users WHERE id=?", id)
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
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Birthday)
		fmt.Println(user.Id, user.Username, user.Password, user.Birthday)
	}
}

func get_user_by_username(username string) User {
	//open database
	db, err := sql.Open("sqlite3", PATH)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	//select user
	var user User
	err = db.QueryRow("SELECT * FROM users WHERE username=?", username).Scan(&user.Id, &user.Username, &user.Password, &user.Birthday)
	if err == sql.ErrNoRows {
		panic(err)
	} else if err != nil {
		panic(err)
	}

	return user

}

func get_user_by_password(password string) User {
	//open database
	db, err := sql.Open("sqlite3", PATH)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	//select user
	var user User
	err = db.QueryRow("SELECT * FROM users WHERE password=?", password).Scan(&user.Id, &user.Username, &user.Password, &user.Birthday)
	if err == sql.ErrNoRows {
		panic(err)
	} else if err != nil {
		panic(err)
	}

	return user

}

func get_user_by_birthday(birthday string) User {
	//open database
	db, err := sql.Open("sqlite3", PATH)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	//select user
	var user User
	err = db.QueryRow("SELECT * FROM users WHERE birthday=?", birthday).Scan(&user.Id, &user.Username, &user.Password, &user.Birthday)
	if err == sql.ErrNoRows {
		panic(err)
	} else if err != nil {
		panic(err)
	}

	return user

}

func get_users_over_age(age int) []User {
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
	var users []User
	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Birthday)
		users = append(users, user)
	}

	return users
}


func main() {
	clear_table()
	create_table()
	add_user("root", "root", "1999-01-01")
	add_user("admin", "admin", "1999-01-02")
	add_user("kiera", "kiera123", "2001-08-12")

	//print users over age 21
	fmt.Println(get_users_over_age(21))

	show_users(PATH)
}


