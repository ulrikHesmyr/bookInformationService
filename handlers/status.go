package handlers

import (
	"assignment-1/utils"
	"encoding/json"
	"net/http"
	"time"
)

/*
The request handler function for the /librarystats/v1/status endpoint
*/
func StatusHandler(duration *time.Time, w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		break
	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	resGutendex, _ := utils.SendGetRequest(utils.GUTENDEX_IP)
	resLanguage, _ := utils.SendGetRequest(utils.LANG2COUNTRY_IP)
	resCountries, _ := utils.SendGetRequest(utils.COUNTRIES_IP + "all")

	uptime := time.Since(*duration).Seconds()

	status_response := utils.StatusResponse{
		Uptime:       uptime,
		Version:      "v1",
		GutendexApi:  resGutendex.StatusCode,
		LanguageApi:  resLanguage.StatusCode,
		CountriesApi: resCountries.StatusCode,
	}

	encoder := json.NewEncoder(w)
	f_err := encoder.Encode(status_response)

	if f_err != nil {
		http.Error(w, "Formatting of data failed", http.StatusInternalServerError)
		return
	}
}
