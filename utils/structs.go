package utils

type BookcountData struct {
	Count   uint32 `json:"count"`
	Next    string `json:"next"`
	Results []struct {
		Authors []struct {
			Name string `json:"name"`
		} `json:"authors"`
	} `json:"results"`
}

type BookcountResponse struct {
	Language string  `json:"language"`
	Books    uint32  `json:"books"`
	Authors  uint32  `json:"authors"`
	Fraction float32 `json:"fraction"`
}
