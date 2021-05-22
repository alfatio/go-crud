package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Username string
	Password string
	Email    string
	Age      int
}

type Users []User

func main() {
	u := User{}
	f := reflect.ValueOf(&u).Elem()
	n := f.NumField()

	c := make([]interface{}, n)

	c[0] = &u.Username
	c[1] = &u.Password
	c[2] = &u.Email
	c[3] = &u.Age

	fmt.Printf("%#v \n", u)

	fmt.Printf("%#v \n", c)

}
