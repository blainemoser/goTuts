package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
		`<h1>Blaine is a Moron</h1>
		<h2>For real, fuck that guy!</h2>
		<p>Seriously, don't we all just hate Blaine!?
		<br>He is probably the most pretentious, annoying motherfucker on this planet!</p>`)
	fmt.Fprintf(w, "<p>You %s even add %s</p>", "can", "<strong>variables</strong>")

}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Here is the about page. Blaine is a fucking moron", r.Context())
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about/", aboutHandler)
	http.ListenAndServe(":8081", nil)
}
