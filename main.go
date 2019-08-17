package main

import (
	"bcard/pkg/db"
	"bcard/router"
)

func main() {
	db.Setup()
	router.InitRouter()
}
