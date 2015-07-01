package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	//"fmt"
	"io/ioutil"
)

func main() {
	//Set up configuration

	//Adding a change

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/Route/{route:.*}", func(w http.ResponseWriter, r *http.Request) {
		DecorateWithLog(RouteEndpoint)(w, r)
	})

	router.HandleFunc("/Poll/{route:.*}", func(w http.ResponseWriter, r *http.Request) {
		DecorateWithLog(Poll)(w, r)
	})

	log.Fatal(http.ListenAndServe(":3000", router))
}

func Poll(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	route := vars["route"]

}

func RouteEndpoint(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	route := vars["route"]

	requestBody = r.Body.Read(p)

	resp, err := http.Get("http://" + route)

	if err == nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {

			_, err = w.Write(body)
			if err == nil {
			}
		}

	}

	return err
}

type appHandler func(http.ResponseWriter, *http.Request) error

func DecorateWithLog(fn appHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			http.Error(w, err.Error(), 500)
		}
	}
}
