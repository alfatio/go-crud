package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPW(str string) (string, error) {

	cost := 6

	h, err := bcrypt.GenerateFromPassword([]byte(str), cost)

	return string(h), err
}

func ComparePW(hashed string, pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pw)) == nil
}
