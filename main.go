package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

const apiKey = "secret_here"
func main() {
	//load in secrets from txt file
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
	// Fatal kills process, good in k8s (spin up new pod)
}

func getList(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Available Endpoints:\n -/weather/{location}\n  -/astronomy/{location}\n -/timezone/{location}\n  -/sports/{location}")
}

func getCurrentWeather(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	location := vars["location"]

	url := "https://weatherapi-com.p.rapidapi.com/current.json?q=" + location
	req, err := http.NewRequest("GET", url, nil) // _ is the error, don't drop it
	// if err != nil
	// kube and io injection
	req.Header.Add("X-RapidAPI-Key", apiKey )
	req.Header.Add("X-RapidAPI-Host", "weatherapi-com.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	// handle error dummy
	defer res.Body.Close() //go read about defer
	body, err := ioutil.ReadAll(res.Body)
	// errors, damnit
	fmt.Fprintf(w, string(body))
}

func getCurrentAstronomyData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	location := vars["location"]

	url := "https://weatherapi-com.p.rapidapi.com/astronomy.json?q=" + location
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", apiKey)
	req.Header.Add("X-RapidAPI-Host", "weatherapi-com.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Fprintf(w, string(body))
}

func getTimeZone(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	location := vars["location"]
	
	url := "https://weatherapi-com.p.rapidapi.com/timezone.json?q=" + location
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-KEY", apiKey)
	req.Header.Add("X-RapidAPI-HOST", "weatherapi-com.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Fprintf(w, string(body))
}

func getSports(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	location := vars["location"]
	
	url := "https://weatherapi-com.p.rapidapi.com/sports.json?q=" + location
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-KEY", apiKey)
	req.Header.Add("X-RapidAPI-HOST", "weatherapi-com.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Fprintf(w, string(body))
}
