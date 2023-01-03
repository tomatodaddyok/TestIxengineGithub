package main

import (
	"log"
	"net/http"

	"github.com/igm/pubsub"
	"github.com/igm/sockjs-go/v3/sockjs"
)

var chat pubsub.Publisher

func main() {
	http.Handle("/echo/", sockjs.NewHandler("/echo", sockjs.DefaultOptions, echoHandler))
	http.Handle("/", http.FileServer(http.Dir("web/")))
	log.Println("Server started on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}