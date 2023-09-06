package main

import (
	"fmt"
	"github.com/gorilla/mux"
	goconfig "github.com/iglin/go-config"
	"io/ioutil"
	"log"
	"net/http"
)

var config = goconfig.NewConfig("./secretConfig.yaml", goconfig.Yaml)
var apiHost = config.GetString("data.apiHost")
var apiKey = config.GetString("data.apiKey")

func main() {	
	handleRequests()
}

func handleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/", getList)
	r.HandleFunc("/weather/{location}", getCurrentWeather)
	r.HandleFunc("/astronomy/{location}", getCurrentAstronomyData)
	r.HandleFunc("/timezone/{location}", getTimeZone)
	r.HandleFunc("/sports/{location}", getSports)
	// bonus endpoint with all four, execute concurrent
	log.Fatal(http.ListenAndServe(":8089", r)) // read about Go Contexts
}

func getList(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Available Endpoints:\n -/weather/{location}\n  -/astronomy/{location}\n -/timezone/{location}\n  -/sports/{location}")
}

func getCurrentWeather(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	location := vars["location"]

	url := "https://weatherapi-com.p.rapidapi.com/current.json?q=" + location
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
	   log.Fatal(err)
	}
	req.Header.Add("X-RapidAPI-Key", apiKey )
	req.Header.Add("X-RapidAPI-Host", apiHost)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
	   log.Fatal(err)
	}

	defer res.Body.Close() //go read about defer
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
	   log.Fatal(err)
	}

	fmt.Fprintf(w, string(body))
}

func getCurrentAstronomyData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	location := vars["location"]

	url := "https://weatherapi-com.p.rapidapi.com/astronomy.json?q=" + location
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
	   log.Fatal(err)
	}

	req.Header.Add("X-RapidAPI-Key", apiKey)
	req.Header.Add("X-RapidAPI-Host", apiHost)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
	   log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
	   log.Fatal(err)
	}

	fmt.Fprintf(w, string(body))
}

func getTimeZone(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	location := vars["location"]
	
	url := "https://weatherapi-com.p.rapidapi.com/timezone.json?q=" + location
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
	   log.Fatal(err)
	}

	req.Header.Add("X-RapidAPI-KEY", apiKey)
	req.Header.Add("X-RapidAPI-HOST", apiHost)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
	   log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
	   log.Fatal(err)
	}

	fmt.Fprintf(w, string(body))
}

func getSports(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	location := vars["location"]
	
	url := "https://weatherapi-com.p.rapidapi.com/sports.json?q=" + location
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
	   log.Fatal(err)
	}

	req.Header.Add("X-RapidAPI-KEY", apiKey)
	req.Header.Add("X-RapidAPI-HOST", apiHost)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
	   log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
	   log.Fatal(err)
	}	

	fmt.Fprintf(w, string(body))
}
