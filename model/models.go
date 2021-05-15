package model

import (
	"log"

	"github.com/alfatio/login/config"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func GetAllUsers() []User {
	db := config.DB()
	query := "SELECT * FROM users"
	var output []User

	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var u User

		err := rows.Scan(&u.Id, &u.Username, &u.Password, &u.Email)
		if err != nil {
			panic(err)
		}
		output = append(output, u)
	}

	return output

}

func GetUserByUsername(u string) User {
	db := config.DB()
	var output User

	query := `
	SELECT * FROM users WHERE username = $1 LIMIT 1
	`
	rows, err := db.Query(query, u)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var u User

		err := rows.Scan(&u.Id, &u.Username, &u.Password, &u.Email)

		if err != nil {
			panic(err)
		}
		output = u
	}

	return output

}

func InsertUser(p User) bool {
	db := config.DB()
	query := `
	INSERT INTO users (username, password, email)
		VALUES ($1, $2, $3)
	`
	_, err := db.Query(query, p.Username, p.Password, p.Email)

	log.Println(err)

	return err == nil
}