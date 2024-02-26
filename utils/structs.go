package utils

type BookcountData struct {
	Count   uint32 `json:"count"`
	Next    string `json:"next"`
	Results []struct {
		Authors []struct {
			Name string `json:"name"`
		} `json:"authors"`
	} `json:"results"`
	AmountAuthors uint32
}

type BookcountResponse struct {
	Language string  `json:"language"`
	Books    uint32  `json:"books"`
	Authors  uint32  `json:"authors"`
	Fraction float32 `json:"fraction"`
}

type Population struct {
	Amount uint32 `json:"population"`
}

type ReadershipResponse struct {
	Country    string `json:"country"`
	Isocode    string `json:"isocode"`
	Books      uint32 `json:"books"`
	Authors    uint32 `json:"authors"`
	Readership uint32 `json:"readership"`
}

type CountryInfo struct {
	Country string `json:"Official_name"`
	Code    string `json:"ISO3166_1_Alpha_2"`
}

type PopulationInfo struct {
	Population uint32 `json:"population"`
}

type StatusResponse struct {
	GutendexApi  int     `json:"gutendexapi"`
	LanguageApi  int     `json:"languageapi"`
	CountriesApi int     `json:"countriesapi"`
	Version      string  `json:"version"`
	Uptime       float64 `json:"uptime"`
}
