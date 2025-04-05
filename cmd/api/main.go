package main

import (
	"fmt"
	"github.com/dotenv-org/godotenvvault"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var Port = ":8080"

func main() {
	err := godotenvvault.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/images/{id}", imageDeleteHandler).Methods("DELETE")
	r.HandleFunc("/api/v1/images", imagesHandler).Methods("GET")

	r.HandleFunc("/api/v1/convert", convertImage).Methods("POST")
	err = http.ListenAndServe(Port, r)
	if err != nil {
		fmt.Println("error while running server ", err)
	}
}
