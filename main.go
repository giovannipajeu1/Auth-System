package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"main.go/configs"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/CreateUser", configs.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", configs.ReadUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", configs.ReadUser).Methods(http.MethodGet)

	fmt.Println("Servidor na Porta 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
