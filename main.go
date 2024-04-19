package main

import (
	"go-basic-crud/config"
	"go-basic-crud/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// 1. Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	log.Println("Starting server on :9003")
	http.ListenAndServe(":9003", nil)
}
