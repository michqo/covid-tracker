package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/gocolly/colly"
)

type Country struct {
	Cases string
	Deaths string
	Recovered string
}

func searchCountry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	c := colly.NewCollector()

	var cases string
	var deaths string
	var recovered string

	c.OnHTML(".content-inner > div:nth-child(6) > div:nth-child(2) > span:nth-child(1)", func(e *colly.HTMLElement) {
		child := e.DOM.First()
		childText := child.Text()
		cases = childText
	})
	c.OnHTML(".content-inner > div:nth-child(7) > div:nth-child(2) > span:nth-child(1)", func(e *colly.HTMLElement) {
		child := e.DOM.First()
		childText := child.Text()
		deaths = childText
	})
	c.OnHTML(".content-inner > div:nth-child(8) > div:nth-child(2) > span:nth-child(1)", func(e *colly.HTMLElement) {
		child := e.DOM.First()
		childText := child.Text()
		recovered = childText
	})
	c.Visit("https://www.worldometers.info/coronavirus/country/" + vars["country"])

	country := Country{ Cases: cases, Deaths: deaths, Recovered: recovered }
	json.NewEncoder(w).Encode(country)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/{country}", searchCountry).Methods("GET")

	fmt.Println("Listening on port 5000")

	handler := cors.Default().Handler(r)
	http.ListenAndServe(":5000", handler)
}
