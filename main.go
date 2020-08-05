package main

import (
	"log"
	"net/http"

	boom "github.com/darahayes/go-boom"
	authentication "github.com/carloserocha/history-application/authentication"
	scavenger "github.com/carloserocha/history-application/scavenger"
	"github.com/gorilla/mux"
)

func main() {
	// fmt.Println(history.SortDescendingAlphabet("sandbox playground"))

	// models.Pool()

	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/scavenger", scavenger.CreateScavenger).Methods("POST")
	router.HandleFunc("/login", authentication.AuthenticateLogin).Methods("POST")

	router.Use(boom.RecoverHandler)

	log.Fatal(http.ListenAndServe(":8001", router))
}
