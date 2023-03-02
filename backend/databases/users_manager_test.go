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

