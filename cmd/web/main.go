package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jordanreger/wx/bots/social"
	"github.com/jordanreger/wx/products/warnings"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		page := `<link rel="stylesheet" href="https://jordanreger.com/lib/style.css"><ul><li><a href="/tornado"><code>tornado</code></a></li><li><a href="/severetstorm"><code>severetstorm</code></a></li><li><a href="/flashflood"><code>flashflood</code></a></li></ul>`
		w.Header().Add("Content-type", "text/html")
		fmt.Fprint(w, page)
	})

	mux.HandleFunc("/tornado", func(w http.ResponseWriter, r *http.Request) {
		warning := warnings.Latest("Tornado warning")
		page := social.GetPost(warning)
		fmt.Fprint(w, page)
	})
	mux.HandleFunc("/severetstorm", func(w http.ResponseWriter, r *http.Request) {
		warning := warnings.Latest("Severe thunderstorm warning")
		page := social.GetPost(warning)
		fmt.Fprint(w, page)
	})
	mux.HandleFunc("/flashflood", func(w http.ResponseWriter, r *http.Request) {
		warning := warnings.Latest("Flash flood warning")
		page := social.GetPost(warning)
		fmt.Fprint(w, page)
	})

	log.Fatal(http.ListenAndServe(":80", mux))
}
