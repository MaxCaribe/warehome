package main

import (
	"github.com/MaxCaribe/warehome/application"
	"log"
	"net/http"
)

func main() {
	server := application.NewServer()
	server.Run()
	err := http.ListenAndServe(":8080", server.Mux)
	if err != nil {
		log.Fatal(err)
	}
}
