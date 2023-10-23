package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	srv := &http.Server{
		Handler:	r,
		Addr:		":8089",
		ReadTimeout:	10 * time.Second,
		WriteTimeout:	10 * time.Second,
	}

	go func() {
	    log.Println("Starting server...")
	    if err := srv.ListenAndServe; err != nil {
	        log.Fatal(err)
	    }
	}()
	
	waitForShutdown(srv)
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

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-interruptChan

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}
