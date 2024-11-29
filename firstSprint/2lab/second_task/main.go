package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Active bool
}

func NewUser(Name string, Age int) *User {
	return &User{Name: Name, Age: Age, Active: true}
}

func main() {
	user := NewUser("Тестовый пользователь", 35)
	fmt.Println(user)
}
