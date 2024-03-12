package main

import (
	"crud/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/users", servidor.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", servidor.SearchUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", servidor.SearchUser).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", servidor.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", servidor.DeleteUser).Methods(http.MethodDelete)

	fmt.Println("server on port 3333")
	log.Fatal(http.ListenAndServe(":3333", router))
}
