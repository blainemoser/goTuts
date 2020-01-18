package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("gets to this point")
	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
}

type newsAggPage struct {
	Title string
	News  string
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := newsAggPage{Title: strings.Title("Missing monkeys found shaved"), News: "Yeah, you read that right"}
	t, _ := template.ParseFiles("pages/basictemplating.html")
	err := t.Execute(w, p)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/get-news", newsAggHandler)
	http.ListenAndServe(":8080", nil)
}
