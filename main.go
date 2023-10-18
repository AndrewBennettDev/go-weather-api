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
	r.HandleFunc("/{location}", getData)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"http://127.0.0.1:3000"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":8089", handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}

func getList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Use /{location} endpoint to get transformed data from Current Weather and Astronomy")
}

func getData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	location := vars["location"]

	url1 := "https://weatherapi-com.p.rapidapi.com/current.json?q=" + location
	req1, err := http.NewRequest("GET", url1, nil)
	if err != nil {
		log.Fatal(err)
	}

	url2 := "https://weatherapi-com.p.rapidapi.com/astronomy.json?q=" + location
	req2, err := http.NewRequest("GET", url2, nil)
	if err != nil {
		log.Fatal(err)
	}

	if apiKey == "secret" {
		fmt.Fprintf(w, "%s", "You need an API key to call this endpoint!")
	} else {

	  req1.Header.Add("X-RapidAPI-Key", apiKey)
	  req1.Header.Add("X-RapidAPI-Host", apiHost)

	  req2.Header.Add("X-RapidAPI-Key", apiKey)
	  req2.Header.Add("X-RapidAPI-Host", apiHost)
	  
	  res1, err := myClient.Do(req1)
	  if err != nil {
		log.Fatal(err)
	  }

	  res2, err := myClient.Do(req2)
	  if err != nil {
		log.Fatal(err)
	  }

	  defer res1.Body.Close()
	  defer res2.Body.Close()

	  weather := new(InputData)
	  astro := new(AstroData)

	  err = json.NewDecoder(res1.Body).Decode(weather)
	  if err != nil {
		log.Fatal(err)
	  }

	  err = json.NewDecoder(res2.Body).Decode(astro)
	  if err != nil {
		log.Fatal(err)
	  }

	  transformedBody := Transform(weather, astro)
	
	  fmt.Fprintf(w, "%+v", transformedBody)
	}
}
