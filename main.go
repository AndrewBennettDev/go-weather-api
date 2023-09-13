package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	goconfig "github.com/iglin/go-config"
	"log"
	"net/http"
	"time"
)

var config = goconfig.NewConfig("./secretConfig.yaml", goconfig.Yaml)
var apiHost = config.GetString("data.apiHost")
var apiKey = config.GetString("data.apiKey")
var myClient = &http.Client{Timeout: 10 * time.Second}

func main() {	
	handleRequests()
}

func handleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/", getList)
	r.HandleFunc("/{endpoint}/{location}", getData)
	log.Fatal(http.ListenAndServe(":8089", r))
}

func getList(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Available Endpoints:\n -/current/{location}\n -/astronomy/{location}\n -/timezone/{location}\n -/sports/{location}")
}

func getData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	endpoint := vars["endpoint"]
	location := vars["location"]
	
	url := "https://weatherapi-com.p.rapidapi.com/" + endpoint + ".json?q=" + location
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
	   log.Fatal(err)
	}
	req.Header.Add("X-RapidAPI-Key", apiKey )
	req.Header.Add("X-RapidAPI-Host", apiHost)

	res, err := myClient.Do(req)
	if err != nil {
	   log.Fatal(err)
	}

	defer res.Body.Close()
	
	target := new(InputData)
	err = json.NewDecoder(res.Body).Decode(target)
	if err != nil {
		log.Fatal(err)
	}

	transformedBody := Transform(target)

	fmt.Fprintf(w, "%+v", transformedBody)
}
