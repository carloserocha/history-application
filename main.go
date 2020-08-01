package main

import (
	"fmt"
	"log"
	"net/http"

	history "github.com/carloserocha/history-application/histories"
	scavenger "github.com/carloserocha/history-application/scavenger"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println(history.SortDescendingAlphabet("sandbox playground"))

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/scavenger", scavenger.CreateScavenger).Methods("POST")

	log.Fatal(http.ListenAndServe(":8001", router))
}
