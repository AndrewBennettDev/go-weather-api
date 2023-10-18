package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	goconfig "github.com/iglin/go-config"
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
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"http://127.0.0.1:3000"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":8089", handlers.CORS(originsOk, headersOk, methodsOk)(r)))
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

	if apiKey == "secret" {
		fmt.Fprintf(w, "%s", "You need an API key to call this endpoint!")
	} else {
		req.Header.Add("X-RapidAPI-Key", apiKey)
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

		transformedBody := Transform(target) // NOTE: this only works for /current/ right now
		
		fmt.Fprintf(w, "%+v", transformedBody)
	}
}
