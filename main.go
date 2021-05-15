package main

import (
	"github.com/alfatio/login/router"

	_ "github.com/lib/pq"
)

func main() {

	r := router.MainRouter()

	r.Run(":3001")
}
