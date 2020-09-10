package main

import (
	"Slack_SlashCommand_Sample/handler"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// connection testAPI
	r.HandleFunc("/ping", handler.Hello).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}
	log.Println("server start!")
	log.Fatal(srv.ListenAndServe())

}
