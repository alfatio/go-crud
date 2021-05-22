package model

import (
	"log"

	"github.com/alfatio/login/config"
	"github.com/alfatio/login/helper"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

var db = config.DB()

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func UserCol(colName string, user *User) interface{} {
	switch colName {
	case "user_id":
		return &user.Id
	case "username":
		return &user.Username
	case "password":
		return &user.Password
	case "email":
		return &user.Email
	default:
		panic("unkown column name " + colName)
	}
}

func GetAllUsers() []User {
	query := "SELECT * FROM users"
	var output []User

	rows, err := db.Query(query)
	panicOnErr(err)

	colNames, err := rows.Columns()
	panicOnErr(err)

	colNum := len(colNames)

	defer rows.Close()

	for rows.Next() {
		var u User

		cols := make([]interface{}, colNum)

		for i := 0; i < colNum; i++ {
			cols[i] = UserCol(colNames[i], &u)

		}

		err := rows.Scan(cols...)

		panicOnErr(err)

		output = append(output, u)
	}

	return output

}

func GetUserByUsername(u string) User {
	var output User

	query := `
	SELECT * FROM users WHERE username = $1 LIMIT 1
	`
	rows, err := db.Query(query, u)
	panicOnErr(err)

	colNames, err := rows.Columns()
	panicOnErr(err)

	colNum := len(colNames)

	defer rows.Close()

	for rows.Next() {
		var u User

		cols := make([]interface{}, colNum)
		for i := 0; i < colNum; i++ {
			cols[i] = UserCol(colNames[i], &u)
		}

		err := rows.Scan(cols...)
		panicOnErr(err)

		output = u
	}

	return output

}

func InsertUser(p User) bool {
	h, err := helper.HashPW(p.Password)

	if err != nil {
		return false
	}

	p.Password = h

	query := `
	INSERT INTO users (username, password, email)
		VALUES ($1, $2, $3)
	`
	_, err = db.Query(query, p.Username, p.Password, p.Email)

	log.Println(err)

	return err == nil
}

func EditUser(p User) (res User, err error) {

	var output User
	query := `
		UPDATE users
		SET username = $1,
				password = $2,
				email 	 = $3
		WHERE user_id = $4
    RETURNING *
	`
	log.Println(p)
	rows, err := db.Query(query, p.Username, p.Password, p.Email, p.Id)
	panicOnErr(err)

	colNames, err := rows.Columns()
	panicOnErr(err)

	colNum := len(colNames)

	for rows.Next() {
		var u User

		cols := make([]interface{}, colNum)

		for i := 0; i < colNum; i++ {
			cols[i] = UserCol(colNames[i], &u)

		}

		rows.Scan(cols...)

		output = u
	}

	return output, nil
}
