package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	dh "../dataHandler"
)

func main() {

	http.HandleFunc("/", HomeHandle)
	http.ListenAndServe(":80", nil)
}

func HomeHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "../index.html")
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}

		sliceOfPoint, err := dh.ParseDataFromString(string(body[:]))
		if err != nil {
			fmt.Fprint(w, "Error. Make sure the coordinates of the points are numbers")
			return
		}
		sliceOfPoint.Sort()
		fmt.Fprint(w, sliceOfPoint.ToString())
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
