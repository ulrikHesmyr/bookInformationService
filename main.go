package main

import (
	"assignment-1/handlers"
	"log"
	"net/http"
	"time"
)

func main() {
	//Variable containing the time of server restart used to provide the "uptime" property of the data response for the "/status/" endpoint
	start := time.Now()

	// //Bookcount endpoint request handler functions
	// router.Path("/librarystats/v1/bookcount/").
	// 	Queries("language", "{language}").
	// 	HandlerFunc(handlers.BookcountHandler)

	// router.HandleFunc("/librarystats/v1/bookcount/", handlers.BookcountInfo)

	http.HandleFunc("/librarystats/v1/bookcount/", handlers.BookcountHandler)

	//Readership endpoint request handler function
	http.HandleFunc("/librarystats/v1/readership/", handlers.ReadershipHandler)

	//Status endpoint request handler function
	http.HandleFunc("/librarystats/v1/status/", func(w http.ResponseWriter, r *http.Request) {
		handlers.StatusHandler(&start, w, r)
	})

	//Http server that will use our pre-configured router as the request handler
	server := &http.Server{
		Handler: nil,
		Addr:    "0.0.0.0:8080",
	}
	log.Fatal(server.ListenAndServe())
}
