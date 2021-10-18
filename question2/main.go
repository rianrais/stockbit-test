package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"question2/omdb"
)

func main() {
	router := httprouter.New()

	// Router list
	router.GET("/omdb", omdb.Get)

	// Listening to port 3070
	log.Fatal(http.ListenAndServe(":3070", router))
}
