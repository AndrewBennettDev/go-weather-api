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
    r.HandleFunc("/weather/{zip}", getCurrentWeather)
    log.Fatal(http.ListenAndServe(":8089", r))
}

func getCurrentWeather(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    zip := vars["zip"]

    url := "https://weatherapi-com.p.rapidapi.com/current.json?q=" + zip
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "[secret_value]")
	req.Header.Add("X-RapidAPI-Host", "weatherapi-com.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Fprintf(w, string(body))
}
