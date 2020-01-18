package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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
	Keywords []string
	Titles   []string
	Contents []string
}

// TODO
// /**
// 	Retrieves the name of the article based on the URL
// 	@param string url
// 	@return string

// */
// func getNameFromURL(url string) string {

// }

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("gets to newsAddHandler")
	a := aggregateWP()
	t, err := template.ParseFiles("pages/aggregatepage.html")
	if err != nil {
		fmt.Println(err)
	}
	err = t.Execute(w, a)
	if err != nil {
		fmt.Println(err)
	}
}

func aggregateWP() map[string]newsMap {
	var s sitemapindex
	var n news
	resp, getRequestError := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	if getRequestError != nil {
		fmt.Println(getRequestError)
		os.Exit(3)
	}
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)
	aggregatedNewsMap := make(map[string]newsMap)
	for _, Location := range s.Locations {
		resp, err := http.Get(strings.TrimSpace(Location))
		if err != nil {
			fmt.Println(err)
			continue
		}
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)

		// Store the data in a News Map
		for idx := range n.Keywords {
			aggregatedNewsMap[n.Titles[idx]] = newsMap{Keyword: n.Keywords[idx], Location: n.Locations[idx]}
		}
	}

	return aggregatedNewsMap
	// for idx, data := range aggregatedNewsMap {
	// 	fmt.Println("\n\n", idx)
	// 	fmt.Println("\n", data.Keyword)
	// 	fmt.Println("\n", data.Location)
	// }
}

func main() {
	http.HandleFunc("/", newsAggHandler)
	http.ListenAndServe(":8080", nil)
}
