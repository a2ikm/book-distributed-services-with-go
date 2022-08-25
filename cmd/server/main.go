package main

import (
	"log"

	"github.com/a2ikm/book-distributed-services-with-go/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
