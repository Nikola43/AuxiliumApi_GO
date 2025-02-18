package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	// Open up our database connection.
	//db, err := sql.Open("mysql", "user:password@tcp(hostIP:port)/database")
	db, err := sql.Open("mysql", "auxilium:-Gdeuapmw18@tcp(127.0.0.1:3306)/auxilium")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}

	// Execute the query
	results, err := db.Query("SELECT * FROM Users")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var user User
		// for each row, scan the result into our tag composite object
		err = results.Scan(&user.ID, &user.LastName, &user.FirstName, &user.IdentificationDocument, &user.Email, &user.Passwd)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		json.NewEncoder(w).Encode(user)
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["ID"]

	// Open up our database connection.
	//db, err := sql.Open("mysql", "user:password@tcp(hostIP:port)/database")
	db, err := sql.Open("mysql", "auxilium:-Gdeuapmw18@tcp(127.0.0.1:3306)/auxilium")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}

	// Execute the query
	results, err := db.Query("SELECT * FROM Users where ID = "+ID)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var user User
		// for each row, scan the result into our tag composite object
		err = results.Scan(&user.ID, &user.LastName, &user.FirstName, &user.IdentificationDocument, &user.Email, &user.Passwd)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		json.NewEncoder(w).Encode(user)
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()
}

func signupUser(w http.ResponseWriter, r *http.Request) {

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println(err)
	}

	// Open up our database connection.
	//db, err := sql.Open("mysql", "user:password@tcp(hostIP:port)/database")
	db, err := sql.Open("mysql", "auxilium:-Gdeuapmw18@tcp(127.0.0.1:3306)/auxilium")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}

	// Execute the query
	results, err := db.Query("INSERT INTO Users VALUES (0, '"+user.LastName+"', '"+user.FirstName+"', '"+user.IdentificationDocument+"', '"+user.Email+"', '"+user.Passwd+"')")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var user User
		// for each row, scan the result into our tag composite object
		err = results.Scan(&user.ID, &user.LastName, &user.FirstName, &user.IdentificationDocument, &user.Email, &user.Passwd)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		json.NewEncoder(w).Encode(user)
		fmt.Fprintf(w, "User inserted")
	}


	// defer the close till after the main function has finished
	// executing
	defer db.Close()
}
