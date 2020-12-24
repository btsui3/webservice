// package main

// import (
// 	"encoding/json"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"strconv"
// 	"strings"
// )

// // French has french vocab and definitions
// type French struct {
// 	French  string `json:"french"`
// 	Value   string `json:"value"`
// 	VocabID int    `json:"vocabID"`
// }

// var frenchList []French

// func init() {
// 	frenchJSON := `[
// 		{
// 			"french": "devoir",
// 			"value": "Must",
// 			"vocabID": 1

// 		},
// 		{
// 			"french": "prendre",
// 			"value": "Take",
// 			"vocabID": 2
// 		},
// 		{
// 			"french": "Pouvoir",
// 			"value": "Can / Able",
// 			"vocabID": 3
// 		},
// 		{
// 			"french": "Partir",
// 			"value": "To Leave",
// 			"vocabID":4
// 		},
// 		{
// 			"french": "Savoir",
// 			"value": "To Know",
// 			"vocabID":5
// 		},
// 		{
// 			"french": "Mettre",
// 			"value": "To Put",
// 			"vocabID": 6
// 		},
// 		{
// 			"french": "connaître",
// 			"value": "To Know",
// 			"vocabID": 7
// 		},
// 		{
// 			"french": "connaître",
// 			"value": "To Know",
// 			"vocabID": 8
// 		}
// 	]`

// 	err := json.Unmarshal([]byte(frenchJSON), &frenchList)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// type fooHandler struct {
// 	Message string
// }

// func frenchWordsHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodGet:
// 		colorsJSON, err := json.Marshal(frenchList)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			log.Fatal(err)
// 			return
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write(colorsJSON)

// 	case http.MethodPost:
// 		//Add new color to the list
// 		var newColor French
// 		bodyBytes, err := ioutil.ReadAll(r.Body)
// 		err = json.Unmarshal(bodyBytes, &newColor)
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			log.Fatal(err)
// 		}

// 		frenchList = append(frenchList, newColor)
// 		w.WriteHeader(http.StatusCreated)
// 		return

// 	}

// }

// func frenchWordHandler(w http.ResponseWriter, r *http.Request) {
// 	urlPathSegments := strings.Split(r.URL.Path, "french/")
// 	vocabID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
// 	if err != nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	}
// 	french, listItemIdex := findVocanByID(vocabID)
// 	if french == nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	}

// 	switch r.Method {
// 	case http.MethodGet:
// 		frenchJSON, err := json.Marshal(french)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write(frenchJSON)

// 	case http.MethodPut:
// 		var updatedFrench French
// 		bodyBytes, err := ioutil.ReadAll(r.Body)
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}
// 		err = json.Unmarshal(bodyBytes, &updatedFrench)
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}

// 		if updatedFrench.VocabID != vocabID {
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}

// 		french = &updatedFrench
// 		frenchList[listItemIdex] = *french
// 		w.WriteHeader(http.StatusOK)
// 		return

// 	default:
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 	}

// }
// func findVocanByID(vocabID int) (*French, int) {
// 	for i, french := range frenchList {
// 		if french.VocabID == vocabID {
// 			return &french, i
// 		}
// 	}
// 	return nil, 0
// }
// func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte(f.Message))
// }

// func barHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("bar called"))
// }

// func main() {
// 	http.Handle("/foo", &fooHandler{Message: "foo called "})
// 	http.HandleFunc("/bar", barHandler)
// 	http.HandleFunc("/french", frenchWordsHandler)
// 	http.HandleFunc("/french/", frenchWordHandler)
// 	http.ListenAndServe(":5000", nil)
// }
