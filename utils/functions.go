package utils

import (
	"encoding/json"
	"net/http"
	"sync"
)

func CountBooksAndAuthors(language string) (BookcountData, error) {

	//Initializing our WaitGroup instance to be able to execute functionality concurrently with goroutines
	var wg = sync.WaitGroup{}
	retrieved_data := BookcountData{}

	//Data variables for tracking amount of unique authors
	var authors_names []string
	var amountAuthors uint32 = 0
	amountChan := make(chan uint32)

	//Initializing client and url for requests to the Gutendex API
	url := GUTENDEX_IP + "/books?languages=" + language

	for url != "" {

		//Own made function made in the "sendhandlers.go" file
		res, err := SendGetRequest(url)

		if err != nil {
			return retrieved_data, err
		}

		//Create instance of <struct> to retrieve and compute data
		decoder := json.NewDecoder(res.Body)
		f_err := decoder.Decode(&retrieved_data)

		if f_err != nil {
			return retrieved_data, f_err
		}

		//Count amount of authors
		wg.Add(1)
		go AuthorCounter(&authors_names, retrieved_data, amountChan)
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
	retrieved_data.AmountAuthors = amountAuthors
	return retrieved_data, nil
}

/*
Function used in goroutines that passes data through a channel
*/
func AuthorCounter(currentAuthors *[]string, newAuthors BookcountData, c chan uint32) {

	var uniqueamount uint32 = 0

	//Converting BookcountData.Results into list of strings

	for i := 0; i < len(newAuthors.Results); i++ {
		for j := 0; j < len(newAuthors.Results[i].Authors); j++ {

			currentAuthor := newAuthors.Results[i].Authors[j].Name
			duplicate := false

			//For each value of newdata, we compare for all the values of currentData
			for k := 0; k < len(*currentAuthors); k++ {
				if (*currentAuthors)[k] == currentAuthor {
					duplicate = true
				}
			}

			if !duplicate {
				/*
					If unique (does not exists any duplicates), then we append this value to
					currentData and increase our uniqueamount
				*/
				*currentAuthors = append(*currentAuthors, currentAuthor)
				uniqueamount++
			}
		}
	}

	//At the end, we put this uniqueamount into c: c <- uniqueamount
	c <- uniqueamount
}

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

func GetCountryInfo(language string) ([]CountryInfo, error) {

	country_data := []CountryInfo{}

	//Gets country official name and ISO code
	res, err := SendGetRequest(LANG2COUNTRY_IP + "language2countries/" + language)

	if err != nil {
		return country_data, err
	}

	//Decoding JSON formatted data from the response
	decoder := json.NewDecoder(res.Body)
	f_err := decoder.Decode(&country_data)

	if f_err != nil {
		return country_data, f_err
	}

	return country_data, nil
}

func GetPopulation(country string) (uint32, error) {
	//Gets population for "country"
	res, err := SendGetRequest(COUNTRIES_IP + "alpha/" + country)

	if err != nil {
		return 0, err
	}

	//Data variable that will store the decoded data
	p := []PopulationInfo{}

	//Decoding the data from JSON format
	decoder := json.NewDecoder(res.Body)
	f_err := decoder.Decode(&p)

	if f_err != nil {
		return 0, err
	}

	return p[0].Population, nil
}

func GetTotalAmountBooks() (float32, error) {
	r, err := SendGetRequest(GUTENDEX_IP + "/books")

	if err != nil {
		return 0, err
	}

	data := BookcountData{}

	decoder := json.NewDecoder(r.Body)
	f_err := decoder.Decode(&data)

	if f_err != nil {
		return 0, err
	}

	return float32(data.Count), nil

}
