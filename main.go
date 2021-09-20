package main

import (
	"backend/router"
	"fmt"

	_ "github.com/pdrum/swagger-automation/docs"
)

func main() {
	fmt.Println("start")
	// model.Create()
	router.HandleRequests()
}
