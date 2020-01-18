package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strings"
	"time"
)

// type location struct {
// 	Loc string `xml:"loc"`
// }

type siteMapIndex struct {
	// Locations []location `xml:"sitemap"`
	Locations []string `xml:"sitemap>loc"`
}

/**
Parses an instance of Location and returns the string(s) contained therein
linked to the location type
This applies a string function to instances of location
*/
// func (l location) String() string {
// 	return fmt.Sprintf(l.Loc)
// }

type egType struct {
	actualProp    string
	nonActualProp string
}

type eg struct {
	egprop egType
}

// News is an aggregation of the news meta-data
type News struct {
	Locations string `xml:"url>loc"`
	LastMods  string `xml:"url>lastmod"`
}

func (e egType) String() string { return fmt.Sprintf("<Hey %s; %s>", e.actualProp, e.nonActualProp) }

func main() {
	// Make a GET request to the specified site
	// The "_" is a placeholder - not a variable declaration but can keep the place
	// of a variable
	// resp, _ := http.Get("https://www.washingtonpost.com/sitemaps/index.xml")
	resp, err := http.Get("https://www.washingtonpost.com/sitemaps/index.xml")

	// aMap := make(map[string]int)

	// Reads the response body

	if err != nil {
		fmt.Println("error received")
		os.Exit(3)
	}
	// This is byte code
	bytes, _ := ioutil.ReadAll(resp.Body)
	parsed := string(bytes)
	// neg := eg{egprop: egType{actualProp: "this is an example", nonActualProp: "this is a non-example"}}
	if strings.Contains(parsed, "404") && strings.Contains(strings.ToLower(parsed), "not found") {
		fmt.Println("Got a 404 here...")
	} else {
		var s siteMapIndex
		var n News
		xml.Unmarshal(bytes, &s)
		// Iterate over the retrived Locations
		// This uses the "range" function - gives the renge (count) of s's Locations array
		// Note that "range" returns the index position (using a "_" placeholder here) and the value
		// This will loop through the categories.
		// Assuming that the articles are not sorted by date; we will get the loop to rank each
		now := time.Now()
		latest := make(map[string]map[string]string)
		for _, Location := range s.Locations {
			// Throws an error without this - trim the leading and trailing white-space from the string
			Location = strings.TrimSpace(Location)
			// fmt.Printf("%s\n", Location)
			resp, err := http.Get(Location)
			if err != nil {
				log.Fatalf("process finished with error = %v", err)
			}
			// fmt.Println(resp)
			latest[Location] = make(map[string]string)
			bytes, _ := ioutil.ReadAll(resp.Body)
			xml.Unmarshal(bytes, &n)
			if len(latest[Location]) == 0 {
				latest[Location][n.LastMods] = n.Locations
			} else {
				for date, loc := range latest[Location] {
					compare, errorDate := time.Parse(time.RFC3339, date)
					current, errorCurrent := time.Parse(time.RFC3339, n.LastMods)
					if errorDate != nil && errorCurrent != nil {
						prev := now.Sub(compare)
						thisone := now.Sub(current)
						if math.Abs(thisone.Minutes()) < math.Abs(prev.Minutes()) {
							latest[Location][date] = loc
						}
					} else {
						log.Println("date error... moving on")
					}
				}
			}
		}

		for url, value := range latest {
			// Pick out the latest ones
			fmt.Printf("URL/Sector: %s\n", url)
			for date, latestStory := range value {
				fmt.Printf("%s\n(%s)\n\n", latestStory, date)
			}
		}
	}
	resp.Body.Close() // Closes the response Body to free resources
}
