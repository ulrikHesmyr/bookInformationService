package handlers

import (
	"assignment-1/utils"
	"encoding/json"
	"net/http"
)

func SendGetRequest(url string) (*http.Response, error) {
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	defer client.CloseIdleConnections()

	res, err := client.Do(r)

	if err != nil {
		return nil, err
	}

	return res, nil

}

func GetTotalAmountBooks() (float32, error) {
	r, err := SendGetRequest(utils.GUTENDEX_IP + "/books")

	if err != nil {
		return 0, err
	}

	data := utils.BookcountData{}

	decoder := json.NewDecoder(r.Body)
	f_err := decoder.Decode(&data)

	if f_err != nil {
		return 0, err
	}

	return float32(data.Count), nil

}
