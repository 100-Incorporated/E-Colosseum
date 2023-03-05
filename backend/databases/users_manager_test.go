package main

import (
	"testing"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func TestAddUser(t *testing.T) {
	//open database
	db, err := sql.Open("sqlite3", PATH)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//insert user

	want := "test test test"
	got := add_user("test", "test", "test")

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}

func TestGetUser(t *testing.T){
	//open database
	db, err := sql.Open("sqlite3", PATH)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//insert user

	want := "test test test"
	got := get_user(4)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGetUserByUsername(t *testing.T){
	//open database
	db, err := sql.Open("sqlite3", PATH)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//insert user

	want := "test"
	got := get_user_by_username("test")

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGetUserByPassword(t *testing.T){
	//open database
	db, err := sql.Open("sqlite3", PATH)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//insert user

	want := "tester tester123 birthday"
	got := get_user_by_password("tester123")

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGetUserByBirthday(t *testing.T){
	//open database
	db, err := sql.Open("sqlite3", PATH)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//insert user

	want := "root root 1999-01-01"
	got := get_user_by_birthday("1999-01-01")

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}