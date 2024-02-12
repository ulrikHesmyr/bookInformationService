package handlers

import (
	"assignment-1/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/gorilla/mux"
)

func BookcountInfo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")

	w.Write([]byte("This is how to use /bookcount endpoint"))
}

func ReadershipInfo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")

	w.Write([]byte("This is how to use the /readership endpoint"))
}

/*
The request handler function for the /librarystats/v1/bookcount endpoint

Base-endpoint must have reader-usable guidance on how to invoke this service
*/
func BookcountHandler(w http.ResponseWriter, r *http.Request) {

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

	//Retrieving all language codes from the query and converting it to a list of strings
	vars := strings.Split(mux.Vars(r)["language"], ",")

	var amount_languages int = len(vars)

	if mux.Vars(r)["language"] == "" {
		http.Error(w, "Misses arguments for the 'language' query", http.StatusBadRequest)
		return
	}

	var data []utils.BookcountResponse

	//Looping through all the requested country-codes to retrieve them as separate objects
	for i := 0; i < amount_languages; i++ {

		//Initializing our WaitGroup instance to be able to execute functionality concurrently with goroutines
		var wg = sync.WaitGroup{}

		//Initializing an instance of a "BookcountRepsons" which will be data retrieved to the client
		response_data := utils.BookcountResponse{Language: vars[i]}

		//Initializing an instance of "BookcountData" which will contain data from the Gutendex API
		retrieved_data := utils.BookcountData{}

		//Data variables for tracking amount of unique authors
		var authors []string
		var amountAuthors uint32 = 0
		amountChan := make(chan uint32)

		//Initializing client and url for requests to the Gutendex API
		url := utils.GUTENDEX_IP + "/books?languages=" + vars[i]

		for url != "" {

			//Own made function made in the "sendhandlers.go" file
			res, err := SendGetRequest(url)

			if err != nil {
				http.Error(w, "Something went wrong when communicating with a third party service", http.StatusBadGateway)
				return
			}

			//Create instance of <struct> to retrieve and compute data
			decoder := json.NewDecoder(res.Body)
			f_err := decoder.Decode(&retrieved_data)

			if f_err != nil {
				http.Error(w, "Something went wrong when formatting data", http.StatusInternalServerError)
				return
			}

			//Checks if there are any result for the requested language code at all
			// if retrieved_data.Count == 0 {
			// 	http.Error(w, "There are not any books for this language code: "+vars[i], http.StatusBadRequest)
			// 	url = ""
			// 	break
			// }

			//Count amount of authors
			wg.Add(1)
			go utils.AuthorCounter(&authors, retrieved_data, amountChan)
			wg.Done()

			x := <-amountChan
			amountAuthors += x

			//Loop termination when last page is reached
			if retrieved_data.Next != url {
				url = retrieved_data.Next
			} else {
				url = ""
			}

		}

		wg.Wait()

		//Set fraction to totalAmountOfBooks / retrieved_data.Count
		f, err := GetTotalAmountBooks()

		if err != nil {
			http.Error(w, "Something went wrong when communicating with a third party service", http.StatusBadGateway)
		}

		numerator := float32(retrieved_data.Count)
		response_data.Fraction = numerator / f

		//Setting amount of books
		response_data.Books = retrieved_data.Count

		//Setting amount of authors
		response_data.Authors = amountAuthors

		//Appending new data to "data"
		data = append(data, response_data)

	}

	fmt.Print(data)

	//Send the "data" to the client
	//Set response header to application/json
	//Use JSON encoder to send data
	encoder := json.NewEncoder(w)
	err := encoder.Encode(data)

	if err != nil {
		http.Error(w, "Something went wrong when formatting data", http.StatusInternalServerError)
		return
	}

}

/*
The request handler function for the /librarystats/v1/readership endpoint

Base-endpoint must have reader-usable guidance on how to invoke this service
*/
func ReadershipHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got a request to readershiphandler")
	switch r.Method {
	case http.MethodGet:
		break
	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}

	fmt.Println("Valid method")
}

/*
The request handler function for the /librarystats/v1/status endpoint
*/
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got a request to statushandler")
	switch r.Method {
	case http.MethodGet:
		break
	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}

	fmt.Println("Valid method")
}
