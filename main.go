package main

import (
	"OCG_final_project_BE/api/router"
	"OCG_final_project_BE/system/seed"

	"fmt"

	_ "github.com/pdrum/swagger-automation/docs"
)

func main() {
	fmt.Println("Start")
	seed.CreateSeed()
	// seed.CreateTable()
	router.HandleRequests()

}
