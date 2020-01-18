package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type sitemapindex struct {
	Locations []string `xml:"sitemap>loc"`
}

type news struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type newsMap struct {
	Keyword  string
	Location string
}

type newsAggPage struct {
	Title string
	News  map[string]newsMap
}

var wg sync.WaitGroup

// TODO
// /**
// 	Retrieves the name of the article based on the URL
// 	@param string url
// 	@return string

// */
// func getNameFromURL(url string) string {

// }

func safety() {
	if r := recover(); r != nil {
		fmt.Println("recovered from error...")
	}
	wg.Done()
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Req: %s %s\n", r.Host, r.URL.Path)
	if strings.Contains(r.URL.Path, "favicon") {
		fmt.Println("Skipping 'favicon' request...")
		return
	}
	a := aggregate()
	t, err := template.ParseFiles("pages/aggregatepage.html")
	if err != nil {
		fmt.Println(err)
	}
	err = t.Execute(w, a)
	if err != nil {
		fmt.Println(err)
	}
}

func (s *sitemapindex) Unmarshalxml(b []byte) {
	fmt.Println(*s)
}

func aggregate() newsAggPage {

	start := time.Now()

	var s sitemapindex
	var n news
	resp, getRequestError := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	if getRequestError != nil {
		fmt.Println(getRequestError)
		// Defer will not be run when using os.Exit()
		os.Exit(3)
	}
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)
	aggregatedNewsMap := make(map[string]newsMap)
	for _, Location := range s.Locations {
		wg.Add(1)
		go func(Location string) {
			defer safety()
			resp, err := http.Get(strings.TrimSpace(Location))
			if err != nil {
				fmt.Println(err)
				return
			}
			bytes, _ := ioutil.ReadAll(resp.Body)
			xml.Unmarshal(bytes, &n)

			// Store the data in a News Map
			for idx := range n.Keywords {
				aggregatedNewsMap[n.Titles[idx]] = newsMap{Keyword: n.Keywords[idx], Location: n.Locations[idx]}
			}
			return
		}(Location)
		wg.Wait()
	}

	elapsed := time.Since(start)
	fmt.Printf("Load time (s): %f\n", elapsed.Seconds())

	// Return an instance of the newsAggPage struct
	return newsAggPage{Title: "News Aggregator", News: aggregatedNewsMap}
}

func main() {
	http.HandleFunc("/", newsAggHandler)
	http.ListenAndServe(":8080", nil)
}
