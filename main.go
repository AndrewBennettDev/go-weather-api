package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	handleRequests()
}

func handleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/weather/{location}", getCurrentWeather)
	r.HandleFunc("/astronomy/{location}", getCurrentAstronomyData)
	log.Fatal(http.ListenAndServe(":8089", r))
}

func getCurrentWeather(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	location := vars["location"]

	url := "https://weatherapi-com.p.rapidapi.com/current.json?q=" + location
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "d403938b14mshe8d8263d5d95dfbp1d81b4jsn9b37fb36178f")
	req.Header.Add("X-RapidAPI-Host", "weatherapi-com.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Fprintf(w, string(body))
}

func getCurrentAstronomyData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	location := vars["location"]

	url := "https://weatherapi-com.p.rapidapi.com/astronomy.json?q=" + location
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "d403938b14mshe8d8263d5d95dfbp1d81b4jsn9b37fb36178f")
	req.Header.Add("X-RapidAPI-Host", "weatherapi-com.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Fprintf(w, string(body))
}
