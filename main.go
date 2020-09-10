package main

import (
	"log"
	"net/http"

	"github.com/Takumaron/Slack_SlashCommand_Sample/handler"
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
