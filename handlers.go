package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	users := UsersList{
		User{ID: "1", LastName: "Soares", FirstName: "Paulo"},
		User{ID: "2", LastName: "Barrera", FirstName: "Miguel"},
	}
	fmt.Println("Endpoint Hit: getAllUsers")

	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	//json.NewEncoder(w).Encode(id)
	fmt.Fprintln(w, "getUser:", id)
}