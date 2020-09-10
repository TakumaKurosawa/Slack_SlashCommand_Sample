package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Takumaron/Slack_SlashCommand_Sample/handler"
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
