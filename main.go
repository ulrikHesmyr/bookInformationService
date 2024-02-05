package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Location struct {
	Name     string `json:"name"`
	Country  string `json:"country,omitempty"` //Allow that the country is not present in the json object
	Postcode uint16 `json:"postcode"`
}

func main() {

	s := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(requestHandler),
	}

	log.Fatal(s.ListenAndServe())
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		defaultGetHandler(w, r)
	case http.MethodPost:
		defaultPostHandler(w, r)
	default:
		http.Error(w, "Method not accepted", http.StatusMethodNotAllowed)
	}
}

func defaultGetHandler(w http.ResponseWriter, r *http.Request) {
	loc := Location{Name: "Skien", Postcode: 3721, Country: "Norge"}

	encoder := json.NewEncoder(w)
	encoder.Encode(loc)

	w.Header().Set("Content-Type", "application/json")

}

func defaultPostHandler(w http.ResponseWriter, r *http.Request) {

	//Put into middleware (???)
	if r.Header.Get("content-type") != "application/json" {
		http.Error(w, "Make sure payload is of JSON format", http.StatusBadRequest)
	}

	//Creating a decoder to match the request body
	decoder := json.NewDecoder(r.Body)

	//Empty data, because we will get the data in the r.Body
	location := Location{}

	//Decoding JSON data from request and parsing it into our "location" variable
	err := decoder.Decode(&location)
	if err != nil {
		http.Error(w, "Server failed to decode JSON data", http.StatusBadRequest)
	}

	fmt.Println(location)
}

//func jsonMiddleware(w http.Response)
