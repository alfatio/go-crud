package main

import (
	"encoding/json"
	"fmt"
)

// type obj struct {
// 	bv string
// 	we int
// 	ol string
// }

type User struct {
	Username string
	Password string
	Email    string
}

type Users []User

func main() {
	var arr Users
	a := Users{
		User{
			Username: "test",
			Password: "qwe",
			Email:    "asd",
		},
		User{
			Username: "test2",
			Password: "qwe2",
			Email:    "asd2",
		},
	}
	arr = a
	b, _ := json.Marshal(arr)
	fmt.Println(string(b))
}
