package main

import (
	"assignment-1/handlers"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	start := time.Now()
	router := mux.NewRouter()

	//Bookcount endpoint request handler functions
	router.Path("/librarystats/v1/bookcount/").
		Queries("language", "{language}").
		HandlerFunc(handlers.BookcountHandler)

	router.HandleFunc("/librarystats/v1/bookcount/", handlers.BookcountInfo)

	//Readership endpoint request handler functions
	router.Path("/librarystats/v1/readership/{language}/").
		Queries("limit", "{limit}").
		HandlerFunc(handlers.ReadershipHandler)

	router.HandleFunc("/librarystats/v1/readership/{language}", handlers.ReadershipHandler)

	router.HandleFunc("/librarystats/v1/readership/", handlers.ReadershipInfo)

	//Status endpoint request handler function
	router.HandleFunc("/librarystats/v1/status/", func(w http.ResponseWriter, r *http.Request) {
		handlers.StatusHandler(&start, w, r)
	})

	//Http server that will use our pre-configured router as the request handler
	server := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
	}
	log.Fatal(server.ListenAndServe())
}
