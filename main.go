package main

import (
	"fmt"
	"log"
	"net/http"

	authentication "github.com/carloerocha/history-application/authentication"
	history "github.com/carloserocha/history-application/histories"
	scavenger "github.com/carloserocha/history-application/scavenger"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println(history.SortDescendingAlphabet("sandbox playground"))

	// models.Pool()

	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/scavenger", scavenger.CreateScavenger).Methods("POST")
	router.HandleFunc("/login", authentication.AuthenticateLogin).Methods("POST")

	log.Fatal(http.ListenAndServe(":8001", router))
}
