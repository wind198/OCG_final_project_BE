package main

import (
	"OCG_final_project_BE/api/router"

	"fmt"

	_ "github.com/pdrum/swagger-automation/docs"
)

func main() {
	fmt.Println("Start")
	// seed.CreateSeed()
	// seed.CreateTable()
	router.HandleRequests()

}
