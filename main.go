package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// French has french vocab and definitions
type French struct {
	French string `json:"french"`
	Value  string `json:"value"`
}

var frenchList []French

func init() {
	colorsJSON := `[
		{
			"french": "devoir",
			"value": "Must"
		},
		{
			"french": "prendre",
			"value": "Take"
		},
		{
			"french": "Pouvoir",
			"value": "Can / Able"
		},
		{
			"french": "Partir",
			"value": "To Leave"
		},
		{
			"french": "Savoir",
			"value": "To Know"
		},
		{
			"french": "Mettre",
			"value": "To Put"
		},
		{
			"french": "connaître",
			"value": "To Know"
		}
	]`

	err := json.Unmarshal([]byte(colorsJSON), &frenchList)
	if err != nil {
		log.Fatal(err)
	}
}

type fooHandler struct {
	Message string
}

func frenchHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		colorsJSON, err := json.Marshal(frenchList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(colorsJSON)

	case http.MethodPost:
		//Add new color to the list
		var newColor French
		bodyBytes, err := ioutil.ReadAll(r.Body)
		err = json.Unmarshal(bodyBytes, &newColor)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Fatal(err)
		}

		frenchList = append(frenchList, newColor)
		w.WriteHeader(http.StatusCreated)
		return

	}

}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.Message))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bar called"))
}

func main() {
	http.Handle("/foo", &fooHandler{Message: "foo called "})
	http.HandleFunc("/bar", barHandler)
	http.HandleFunc("/colors", frenchHandler)
	http.ListenAndServe(":5000", nil)
}
