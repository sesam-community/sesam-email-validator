package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

var emailRegex *regexp.Regexp = regexp.MustCompile(`^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting service on port %s", port)

	router := mux.NewRouter()
	//path parameter contains field containing email address
	router.HandleFunc("/{fieldName}", ProcessMessages).Methods("POST")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}

//Service entry point
func ProcessMessages(w http.ResponseWriter, r *http.Request) {
	var bodyJsonArray []interface{}

	params := mux.Vars(r)
	fieldToValidate := params["fieldName"]

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		returnInternalErrorResponse(err, w)
		return
	}

	err = json.Unmarshal(body, &bodyJsonArray)

	if err != nil {
		returnInternalErrorResponse(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("["))
	for index, item := range bodyJsonArray {
		mappedItem := item.(map[string]interface{})
		mappedItem[fieldToValidate+"_validated"] = mappedItem[fieldToValidate] != nil && emailRegex.MatchString(mappedItem[fieldToValidate].(string))
		if index > 0 {
			w.Write([]byte(","))
		}
		jsonData, err := json.Marshal(mappedItem)
		if err != nil {
			log.Printf("error whhile marshalling map to json %s", err.Error())
			continue
		}
		w.Write(jsonData)
	}
	w.Write([]byte("]"))
}

func returnInternalErrorResponse(err error, w http.ResponseWriter) {
	log.Println(err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
