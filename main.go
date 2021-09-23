package main

import (
	"backend/api/router"
	"fmt"

	_ "github.com/pdrum/swagger-automation/docs"
)

func main() {
	fmt.Println("Start")
	// seed.CreateTable()
	router.HandleRequests()

}
