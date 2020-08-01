package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	history "github.com/carloserocha/history-application/histories"
	scavenger "github.com/carloserocha/history-application/scavenger"
)
// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, "+req.URL.Query().Get(":anagram")+"!\n")
}

func main() {
	fmt.Println(history.SortDescendingAlphabet("sandbox playground"))

	router := mux.NewRouter()
	router.HandleFunc("/scavenger/", scavenger.CreateScavenger).Methods("POST")

	log.Fatal(http.ListenAndServe(":8001", router))
}