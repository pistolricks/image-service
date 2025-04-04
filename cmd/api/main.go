package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var Port = ":8080"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/convert", ConvertImage).Methods("POST")
	err := http.ListenAndServe(Port, r)
	if err != nil {
		fmt.Println("error while running server ", err)
	}
}
