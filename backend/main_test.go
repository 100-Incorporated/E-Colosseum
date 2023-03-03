package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Function to make JSON list match the format from c.IndentedJSON
// There is added logic to stop the keys from being alphabetical
func formatJsonList(jsonString string) string {
	var users []UserResponse
	if err := json.Unmarshal([]byte(jsonString), &users); err != nil {
		panic(err)
	}

	// Decode the JSON string into a map while preserving key order
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "    ")
	enc.SetEscapeHTML(false)

	// Sort the users by ID to match c.IndentedJSON
	sort.SliceStable(users, func(i, j int) bool {
		return users[i].ID < users[j].ID
	})

	if err := enc.Encode(users); err != nil {
		panic(err)
	}

	// Removing the extra newline to match c.IndentedJSON
	return strings.TrimRight(buf.String(), "\n")
}

func formatJson(jsonString string) string {
	var user UserResponse
	if err := json.Unmarshal([]byte(jsonString), &user); err != nil {
		panic(err)
	}

	// Decode the JSON string into a map while preserving key order
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "    ")
	enc.SetEscapeHTML(false)

	if err := enc.Encode(user); err != nil {
		panic(err)
	}

	// Removing the extra newline to match c.IndentedJSON
	return strings.TrimRight(buf.String(), "\n")
}

func TestGetUsers(t *testing.T) {
	// Setup
	db, err := sql.Open("sqlite3", "./databases/users.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := setupRouter(db)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(w, req)

	// Check the response body is what we expect.
	expected := `[{"id":1,"username":"root","password":"root","birthday":"1999-01-01"},{"id":2,"username":"admin","password":"admin","birthday":"1999-01-01"}]`

	assert.Equal(t, formatJsonList(expected), w.Body.String(), "Response body mismatch")
}

func TestGetUserByID(t *testing.T) {
	// Setup
	db, err := sql.Open("sqlite3", "./databases/users.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := setupRouter(db)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(w, req)

	// Check the response body is what we expect.
	expected := `{"id":1,"username":"root","password":"root","birthday":"1999-01-01"}`

	assert.Equal(t, formatJson(expected), w.Body.String(), "Response body mismatch")
}

func TestCreateUser(t *testing.T) {
	// Setup
	db, err := sql.Open("sqlite3", "./databases/users.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := setupRouter(db)

	// Defer function to delete the user after the test
	defer func() {
		_, err := db.Exec("DELETE FROM users WHERE id = 3")
		if err != nil {
			panic(err)
		}
	}()

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/users", strings.NewReader(`{"username":"POSTtest","password":"POSTtest","birthday":"1999-01-01"}`))
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(w, req)

	// Check the response body is what we expect.
	expected := `{"id":3,"username":"POSTtest","password":"POSTtest","birthday":"1999-01-01"}`
	assert.Equal(t, formatJson(expected), w.Body.String(), "Response body mismatch")

	// Check the user was actually created
	w = httptest.NewRecorder()
	req, err = http.NewRequest("GET", "/users/3", nil)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(w, req)

	expected = `{"id":3,"username":"POSTtest","password":"POSTtest","birthday":"1999-01-01"}`
	assert.Equal(t, formatJson(expected), w.Body.String(), "Response body mismatch")
}

func TestDeleteUser(t *testing.T) {
	// Setup
	db, err := sql.Open("sqlite3", "./databases/users.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := setupRouter(db)

	// Create a user to delete for the test
	_, err = db.Exec("INSERT INTO users (id, username, password, birthday) VALUES (3, 'DELETEtest', 'DELETEtest', '1999-01-01')")

	// Defer function to delete the user after the test in case the test fails
	defer func() {
		_, err := db.Exec("DELETE FROM users WHERE id = 3")
		if err != nil {
			panic(err)
		}
	}()

	w := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/users/3", nil)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(w, req)

	// I refuse to write a third JSON formatter function
	expected := `{
    "message": "User deleted successfully"
}`
	assert.Equal(t, (expected), w.Body.String(), "Response body mismatch")

	// Check the user was actually deleted
	w = httptest.NewRecorder()
	req, err = http.NewRequest("GET", "/users/3", nil)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(w, req)

	expected = `{
    "message": "User not found"
}`
	assert.Equal(t, (expected), w.Body.String(), "Response body mismatch")
}

func TestPutUser(t *testing.T) {
	// Setup
	db, err := sql.Open("sqlite3", "./databases/users.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := setupRouter(db)

	// Create a user to update for the test
	_, err = db.Exec("INSERT INTO users (id, username, password, birthday) VALUES (3, 'test', 'test', '1999-01-01')")

	// Defer function to delete the user after the test
	defer func() {
		_, err := db.Exec("DELETE FROM users WHERE id = 3")
		if err != nil {
			panic(err)
		}
	}()

	w := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", "/users/3", strings.NewReader(`{"username":"PUTtest","password":"PUTtest","birthday":"1999-01-01"}`))
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(w, req)

	// Check the response body is what we expect.
	expected := `{"id":3,"username":"PUTtest","password":"PUTtest","birthday":"1999-01-01"}`
	assert.Equal(t, formatJson(expected), w.Body.String(), "Response body mismatch")

	// Check the user was actually updated
	w = httptest.NewRecorder()
	req, err = http.NewRequest("GET", "/users/3", nil)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(w, req)

	expected = `{"id":3,"username":"PUTtest","password":"PUTtest","birthday":"1999-01-01"}`
	assert.Equal(t, formatJson(expected), w.Body.String(), "Response body mismatch")
}

func TestPatchUser(t *testing.T) {
	// Setup
	db, err := sql.Open("sqlite3", "./databases/users.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := setupRouter(db)

	// Create a user to update for the test
	_, err = db.Exec("INSERT INTO users (id, username, password, birthday) VALUES (3, 'test', 'test', '1999-01-01')")

	// Defer function to delete the user after the test
	defer func() {
		_, err := db.Exec("DELETE FROM users WHERE id = 3")
		if err != nil {
			panic(err)
		}
	}()

	w := httptest.NewRecorder()
	req, err := http.NewRequest("PATCH", "/users/3", strings.NewReader(`{"username":"PATCHtest"}`))
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(w, req)

	// Check the response body is what we expect.
	expected := `{"id":3,"username":"PATCHtest","password":"test","birthday":"1999-01-01"}`
	assert.Equal(t, formatJson(expected), w.Body.String(), "Response body mismatch")

	// Check the user was actually updated
	w = httptest.NewRecorder()
	req, err = http.NewRequest("GET", "/users/3", nil)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(w, req)

	expected = `{"id":3,"username":"PATCHtest","password":"test","birthday":"1999-01-01"}`
	assert.Equal(t, formatJson(expected), w.Body.String(), "Response body mismatch")

	// Check that the user can be updated with multiple fields
	w = httptest.NewRecorder()
	req, err = http.NewRequest("PATCH", "/users/3", strings.NewReader(`{"username":"PATCHtest2","password":"PATCHtest2"}`))
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(w, req)

	// Check the response body is what we expect.
	expected = `{"id":3,"username":"PATCHtest2","password":"PATCHtest2","birthday":"1999-01-01"}`
	assert.Equal(t, formatJson(expected), w.Body.String(), "Response body mismatch")
}

func TestQueryOutOfRange(t *testing.T) {
	// Setup
	db, err := sql.Open("sqlite3", "./databases/users.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := setupRouter(db)

	// GET /users/id
	// Test that a user that doesn't exist returns a 404
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/users/999", nil)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(w, req)

	expected := `{
    "message": "User not found"
}`
	assert.Equal(t, (expected), w.Body.String(), "Response body mismatch")

	// DELETE /users/id
	// Test that a user that doesn't exist returns a 404
	w = httptest.NewRecorder()
	req, err = http.NewRequest("DELETE", "/users/999", nil)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(w, req)

	expected = `{
    "message": "User not found"
}`
	assert.Equal(t, (expected), w.Body.String(), "Response body mismatch")

	// PUT /users/id
	// Test that a user that doesn't exist returns a 404
	w = httptest.NewRecorder()
	req, err = http.NewRequest("PUT", "/users/999", nil)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(w, req)

	expected = `{
    "message": "User not found"
}`
	assert.Equal(t, (expected), w.Body.String(), "Response body mismatch")

	// PATCH /users/id
	// Test that a user that doesn't exist returns a 404
	w = httptest.NewRecorder()
	req, err = http.NewRequest("PATCH", "/users/999", nil)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(w, req)

	expected = `{
    "message": "User not found"
}`
	assert.Equal(t, (expected), w.Body.String(), "Response body mismatch")
}

func TestInvalidRequestBody(t *testing.T) {
	// Setup
	db, err := sql.Open("sqlite3", "./databases/users.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := setupRouter(db)

	// POST /users
	// Test that an invalid request is handled correctly
	w := httptest.NewRecorder()
	// Testing a request with no body
	req, err := http.NewRequest("POST", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(w, req)

	expected := `{
    "message": "Invalid request body"
}`
	assert.Equal(t, (expected), w.Body.String(), "Response body mismatch")

	// Testing a request with an invalid body
	w = httptest.NewRecorder()
	req, err = http.NewRequest("POST", "/users", strings.NewReader(`{"username":"test"}`))
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(w, req)

	expected = `{
    "message": "Invalid request body"
}`
	assert.Equal(t, (expected), w.Body.String(), "Response body mismatch")

	// Testing a response with invalid JSON format
	w = httptest.NewRecorder()
	// Note that  a " is missing from the end of the username field
	req, err = http.NewRequest("POST", "/users", strings.NewReader(`{"username":"test,"password":"test","birthday":"1999-01-01"}`))
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(w, req)

	expected = `{
    "message": "Invalid request body"
}`
	assert.Equal(t, (expected), w.Body.String(), "Response body mismatch")

	// PUT /users
	// Create a user to update
	w = httptest.NewRecorder()
	req, err = http.NewRequest("POST", "/users", strings.NewReader(`{"username":"test","password":"test","birthday":"1999-01-01"}`))
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(w, req)

	// Defer the deletion of the user in case the test fails
	defer func() {
		w = httptest.NewRecorder()
		req, err = http.NewRequest("DELETE", "/users/3", nil)
		if err != nil {
			t.Fatal(err)
		}
		router.ServeHTTP(w, req)
	}()

	// Testing a request with an incomplete body
	w = httptest.NewRecorder()
	req, err = http.NewRequest("PUT", "/users/3", strings.NewReader(`{"username":"test"}`))
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(w, req)

	expected = `{
    "message": "Invalid request body"
}`
	assert.Equal(t, (expected), w.Body.String(), "Response body mismatch")

	// Testing a request with an invalid body
	w = httptest.NewRecorder()
	// Note that the username field is misspelled
	req, err = http.NewRequest("PUT", "/users/3", strings.NewReader(`{"usernme":"test","password":"test","birthday":"1999-01-01"}`))
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(w, req)

	expected = `{
    "message": "Invalid request body"
}`
	assert.Equal(t, (expected), w.Body.String(), "Response body mismatch")

	// Testing a response with invalid JSON format
	w = httptest.NewRecorder()
	// Note that  a " is missing from the end of the username field
	req, err = http.NewRequest("PUT", "/users/3", strings.NewReader(`{"username":"test,"password":"test","birthday":"1999-01-01"}`))
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(w, req)

	expected = `{
    "message": "Invalid request body"
}`
	assert.Equal(t, (expected), w.Body.String(), "Response body mismatch")

	// PATCH /users
	// Test that an invalid request is handled correctly
	w = httptest.NewRecorder()
	// Testing a request with no body
	req, err = http.NewRequest("PATCH", "/users/3", nil)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(w, req)

	expected = `{
    "message": "Invalid request body"
}`
	assert.Equal(t, (expected), w.Body.String(), "Response body mismatch")

	// Testing a request with an invalid body
	w = httptest.NewRecorder()
	// Note that the username field is misspelled
	req, err = http.NewRequest("PATCH", "/users/3", strings.NewReader(`{"usernme":"test"}`))
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(w, req)

	expected = `{
    "message": "Invalid request body"
}`
	assert.Equal(t, (expected), w.Body.String(), "Response body mismatch")

	// Testing a response with invalid JSON format
	w = httptest.NewRecorder()
	// Note that  a " is missing from the end of the username field
	req, err = http.NewRequest("PATCH", "/users/3", strings.NewReader(`{"username":"test,"password":"test","birthday":"1999-01-01"}`))
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(w, req)

	expected = `{
    "message": "Invalid request body"
}`
	assert.Equal(t, (expected), w.Body.String(), "Response body mismatch")
}
