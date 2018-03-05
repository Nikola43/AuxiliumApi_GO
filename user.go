package main

type User struct {
	ID string `json:"ID"`
	LastName string `json:"LastName"`
	FirstName string `json:"FirstName"`
}

type UsersList []User