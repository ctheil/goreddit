package main

import (
	"log"
	"net/http"

	"github.com/ctheil/goreddit/postgres"
	"github.com/ctheil/goreddit/web"
)

func main() {

	store, err := postgres.NewStore("postgres://postgres:secret@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	h := web.NewHandler(store)

	if err := http.ListenAndServe(":3000", h); err != nil {
		log.Fatal(err)
	}
}
