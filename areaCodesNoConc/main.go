package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

// Struct to map a single country data point
type countryInfo struct {
	Name       string `json:"name"`
	Capital    string `json:"capital"`
	Region     string `json:"region"`
	Subregion  string `json:"subregion"`
	Population int    `json:"population"`
}

type indexPage struct {
	Title string
	Stats map[string]countryInfo
}

// TODO
// /**
// 	Retrieves the name of the article based on the URL
// 	@param string url
// 	@return string

// */
// func getNameFromURL(url string) string {

// }

func getCountriesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Req: %s %s\n", r.Host, r.URL.Path)
	if strings.Contains(r.URL.Path, "favicon") {
		fmt.Println("Skipping 'favicon' request...")
		return
	}
	a := fetch()
	t, err := template.ParseFiles("pages/list.html")
	if err != nil {
		fmt.Println(err)
	}
	err = t.Execute(w, a)
	if err != nil {
		fmt.Println(err)
	}
}

func getCountryDetails(name string) []countryInfo {
	resp, getRequestError := http.Get("https://restcountries.eu/rest/v2/name/" + name)
	if getRequestError != nil {
		fmt.Println(getRequestError)
		os.Exit(3)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("Server made yucky woo!")
	}
	s := make([]countryInfo, 0)
	json.Unmarshal(bytes, &s)
	return s
}

func fetch() indexPage {

	start := time.Now()

	// Create a list of the countries to be looked up
	countryNames := []string{"south africa", "usa", "spain", "china", "funkytown", "uganda", "ukrain", "Trinidad and Tobago", "Togo", "montserat", "turkey", "chile", "argentina", "germany"}
	countries := make(map[string]countryInfo)

	for _, country := range countryNames {
		// go func(country string) {
		details := getCountryDetails(country)
		for _, countryDetails := range details {
			countries[country] = countryDetails
		}
		// }(country)
	}

	elapsed := time.Since(start)
	fmt.Printf("Load time (s): %f\n", elapsed.Seconds())

	// Return an instance of the indexPage struct
	return indexPage{Title: "Countries", Stats: countries}
}

func main() {
	http.HandleFunc("/", getCountriesHandler)
	http.ListenAndServe(":8080", nil)
}
