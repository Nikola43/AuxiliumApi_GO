package main

import "fmt"

type User struct {
	ID                     string    `json:"ID"`
	LastName               string `json:"LastName"`
	FirstName              string `json:"FirstName"`
	IdentificationDocument string `json:"IdentificationDocument"`
	Email                  string `json:"Email"`
	Passwd                 string `json:"Passwd"`
}

func (u *User) showUserInfo() {
	fmt.Println(u.LastName+","+u.FirstName+","+u.IdentificationDocument+","+u.Email+","+u.Passwd)
}