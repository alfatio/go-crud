package config

import (
	"database/sql"
	"fmt"
	"log"

	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func DB() *sql.DB {

	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}
	dbString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", dbString)

	if err != nil {
		panic(err)
	}

	return db
}
