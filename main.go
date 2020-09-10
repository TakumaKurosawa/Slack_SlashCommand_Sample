package main

import (
	"log"
	"net/http"
	"slackSlashCommandApp/handler"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// connection testAPI
	r.HandleFunc("/ping", handler.Hello).Methods("GET")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Println("server start!")
	log.Fatal(srv.ListenAndServe())

}
