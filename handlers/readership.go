package handlers

import (
	"assignment-1/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

/*
Function to retrieve human-readable user guide for the "readership" endpoint
*/
func ReadershipInfo(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/readership.html")
}

/*
The request handler function for the /librarystats/v1/readership endpoint

Base-endpoint must have reader-usable guidance on how to invoke this service
*/
func ReadershipHandler(w http.ResponseWriter, r *http.Request) {

	path := strings.TrimPrefix(r.URL.Path, "/librarystats/v1/readership")

	pathArgs := strings.Split(path, "/")

	amountArgs := len(pathArgs)

	var language string
	var amount int

	if path == "/" {
		ReadershipInfo(w, r)
		return
	} else if amountArgs == 2 || amountArgs == 3 {
		language = pathArgs[1]
	} else {
		http.Error(w, "Request not valid", http.StatusBadRequest)
	}

	//Retrieving data in the JSON format, therefore we specify it in the Headers for the browser to format accordingly
	w.Header().Add("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		break
	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}

	//Check if valid request (limit is a number), else return http.StatusBadRequest

	//	Henter authors og books med utils.CountBooksAndAuthors
	//		Henter properties: AmountAuthors og Count
	c, err := utils.CountBooksAndAuthors(language)
	if err != nil {
		http.Error(w, "Failed to retrieve bookcount and amount of unique authors from a third party service", http.StatusBadGateway)
		return
	}

	countries, err := utils.GetCountryInfo(language)

	if err != nil {
		http.Error(w, "Failed to retrieve country and language-code information from a third party service", http.StatusBadGateway)
		return
	}

	response_data := []utils.ReadershipResponse{}

	if len(r.URL.Query()["limit"]) == 0 {
		amount = len(countries)
	} else {
		amount, err = strconv.Atoi(r.URL.Query()["limit"][0])
		if err != nil {
			http.Error(w, "limit query argument is not a valid number. Must be an integer.", http.StatusBadRequest)
			return
		}
	}

	//Loop that retrieves the population/readership of each country in addition to compute and structure the data that we will retrieve to the client
	for i := 0; i < amount && i < len(countries); i++ {

		populationInfo, err := utils.GetPopulation(countries[i].Code)
		if err != nil {
			http.Error(w, "Failed to retrieve info about population for country-code "+countries[i].Code+" from a third party service", http.StatusBadGateway)
			return
		}

		country := utils.ReadershipResponse{Readership: populationInfo, Authors: c.AmountAuthors, Books: c.Count, Isocode: countries[i].Code, Country: countries[i].Country}
		response_data = append(response_data, country)
	}

	//Encode response_data to JSON to send back to the client
	encoder := json.NewEncoder(w)
	f_err := encoder.Encode(response_data)

	if f_err != nil {
		http.Error(w, "Failed to parse into JSON data format", http.StatusInternalServerError)
	}
}
