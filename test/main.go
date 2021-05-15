package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string
	Password string
	Email    string
}

type Users []User

var pwd string = "$2a$05$xy4PBOiqqMx3CxiwqO2YcOC6GbRxJE/lvWjgZtj3.6TIDgrjnqN.2"

func main() {

	p := "passwor"
	// s := 5

	// hashed, err := bcrypt.GenerateFromPassword([]byte(p), s)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(string(hashed))

	err := bcrypt.CompareHashAndPassword([]byte(pwd), []byte(p))

	fmt.Println(err)
}
