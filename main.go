package main

import (
	"fmt"
	"log"
	"net/http"
	"todolist/cmd/api/config"
	"todolist/cmd/api/router"
)

func main() {
	config.LoadEnvVariables()

	router := router.CreateRouter()

	fmt.Printf("Server running %d", config.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router))
}
