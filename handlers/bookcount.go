package handlers

import (
	"assignment-1/utils"
	"encoding/json"
	"net/http"
	"strings"
)

/*
Function to retrieve a human-readable user guide for the "bookcount" endpoint of our API
*/
func BookcountInfo(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/bookcount.html")
}

/*
The request handler function for the /librarystats/v1/bookcount endpoint

Base-endpoint must have reader-usable guidance on how to invoke this service
*/
func BookcountHandler(w http.ResponseWriter, r *http.Request) {

	//Retrieving all language codes from the query and converting it to a list of strings
	var languages []string
	path := strings.TrimPrefix(r.URL.Path, "/librarystats/v1/bookcount")

	if len(r.URL.Query()["language"]) > 0 {
		languages = strings.Split(r.URL.Query()["language"][0], ",")
	}

	if path == "/" && len(languages) == 0 {
		BookcountInfo(w, r)
		return
	} else if path != "/" {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	//Retrieving data in the JSON format, therefore we specify it in the Headers for the browser to format accordingly
	w.Header().Add("Content-Type", "application/json")

	//Only handling requests that are supported for this route
	switch r.Method {
	case http.MethodGet:
		break
	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}

	var amount_languages int = len(languages)

	var data []utils.BookcountResponse

	//Looping through all the requested country-codes to retrieve them as separate objects
	for i := 0; i < amount_languages; i++ {

		//Initializing an instance of a "BookcountRepsons" which will be data retrieved to the client
		response_data := utils.BookcountResponse{Language: languages[i]}

		//Initializing an instance of "BookcountData" which will contain data from the Gutendex API
		retrieved_data, err := utils.CountBooksAndAuthors(languages[i])

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
		}

		//Set fraction to totalAmountOfBooks / retrieved_data.Count
		f, err := utils.GetTotalAmountBooks()

		if err != nil {
			http.Error(w, "Something went wrong when communicating with a third party service", http.StatusBadGateway)
		}

		//Setting the fraction
		numerator := float32(retrieved_data.Count)
		response_data.Fraction = numerator / f

		//Setting amount of books
		response_data.Books = retrieved_data.Count

		//Setting amount of authors
		response_data.Authors = retrieved_data.AmountAuthors

		//Appending new data to "data"
		data = append(data, response_data)

	}

	//Uses JSON encoder to encode into JSON format and sends data
	encoder := json.NewEncoder(w)
	err := encoder.Encode(data)

	if err != nil {
		http.Error(w, "Something went wrong when formatting data", http.StatusInternalServerError)
		return
	}

}
