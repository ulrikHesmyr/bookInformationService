package handlers

import (
	"fmt"
	"net/http"
)

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
