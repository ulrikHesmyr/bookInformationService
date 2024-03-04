# Usage and setup for the "Book information service" API

## Setup
### Prerequisites
1. Make sure go (Requires go version 1.21.0 or newer) is downloaded and installed on your computer. [See go.dev/doc/install](https://go.dev/doc/install)
2. Check if go is properly installed ``` go version ```. If you get a output specifying the version of go that is installed, we are good to 
go. If it does not display, try to restart your terminal (close and reopen) and run the command once more.

### Installation
1. Clone the repository by using ``` git clone <link to this repository> ```. Now you should see a new folder with the name of the 
git repository
2. Then go to the root folder of the git repository ``` cd <name of repository folder> ```
3. Run ``` go run main.go ```

---

## Usage
This API provides three endpoints with related paths and queries. Beneath you will find documentation for each endpoint containing
info about accessible data, data formats, data structure, arguments and parameter values, presence and example requests.

### Endpoint: /bookcount
The bookcount endpoint is open for GET requests and will retrieve the bookcount for any given language(s). The language is identified via 2-letter ISO country code.
The data 

#### /bookcount paths:
##### /
This path provides a full user-readable documentation for the /bookcount endpoint provided in format of a HTML document. The document includes information of paths, queries, data format, data structure, accepted values as argument to paths and queries and 
sample requests with corresponding response-data.

##### /?language
This query of the default-path "/" takes either a single-value or a comma-separated list [2-letter ISO country-code](https://en.wikipedia.org/wiki/List_of_ISO_639_language_codes). At least a single-value argument to the query is mandatory to retrieve data. For each value of the argument, the list returned by the request will containt a corresponding object. The response data is 
of JSON format (conent type of application and sub-level media type of json). The data structure has the following properties:
- language (string): The language code which is the same as the value of the query
- books (uint32): The amount of books available for the queried language
- authors (uint32): The amount of unique authors
- fraction (float32): The number of books divided by the number of all books served via gutendex

Example:
- Request: ``` /bookcount/?language=hu ```
```
[{ 
    "language": "hu",
    "books": 532, 
    "authors": 158, 
    "fraction":  0.007298369
}]
```

- Request: ``` /bookcount/?language=hu,fi ```
```
[
    {
        "language": "hu",
        "books": 532,
        "authors": 158,
        "fraction": 0.007298369
    },
    {
        "language": "fi",
        "books": 2824,
        "authors": 884,
        "fraction": 0.03874172
    }
]
```


### Endpoint: /readership

#### /readership paths:
##### /
This path provides a full user-readable documentation for the /readership endpoint provided in format of a HTML document. The document includes information of paths, queries, data format, data structure, accepted values as argument to paths and queries and 
sample requests with corresponding response-data.

##### /{language}
This path takes a single-value [2-letter ISO country-code](https://en.wikipedia.org/wiki/List_of_ISO_639_language_codes) Set 1, as argument 
which will return a list with an object for each language spoken in the country represented by the 2-letter ISO country-code. The response data is of JSON format (conent type of application and sub-level media type of json). The data structure has the following properties:
- country (string): The official name of the country
- isocode (string): The 2-letter ISO country-code, corresponding to the one provided in the request
- books (uint32): The amount of books available for the queried language
- authors (uint32): The amount of unique authors
- readerhip (uint32): The population of the country

Example:
- Request: ``` /readership/no ```
```
[
    {
        "country": "Iceland",
        "isocode": "IS",
        "books": 21,
        "authors": 16,
        "readership": 366425
    },
    {
        "country": "Norway",
        "isocode": "NO",
        "books": 21,
        "authors": 16,
        "readership": 5379475
    },
    {
        "country": "Svalbard and Jan Mayen Islands",
        "isocode": "SJ",
        "books": 21,
        "authors": 16,
        "readership": 2562
    }
]
```


##### /{language}/?limit={limit}
The "limit" query of this path is optional, but requires an integer value to specify the maximum amount of objects returned in the list.


Example:
- Request: ``` /readership/fi/?limit=2 ```
```
[
    {
        "country": "Finland",
        "isocode": "FI",
        "books": 2824,
        "authors": 884,
        "readership": 5530719
    },
    {
        "country": "Norway",
        "isocode": "NO",
        "books": 2824,
        "authors": 884,
        "readership": 5379475
    }
]
```

### Endpoint: /status

#### /status paths:
##### /
This path provides a diagnostics interface that indicates the availability of individual third party services that the API depends on to 
provide full functionality. The data has JSON format and contains the following properties:
- gutendexapi (int): The statuscode returned when sending requests to the Gutendex API
- languageapi (int): The statuscode returned when sending requests to the language2countries API
- countriesapi (int): The statuscode returned when sending requests to the restcountries API
- version (string): version of our API
- uptime (float64): Time in seconds from the last service restart

Example: 
- Request: ``` /status/ ```

```
{
    "gutendexapi": 200,
    "languageapi": 200,
    "countriesapi": 200,
    "version": "v1",
    "uptime": 6.0697663
}
```
