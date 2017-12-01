package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type redirectData struct {
	Name     string `json:"Name"`
	Code     int32  `json:"Code"`
	Redirect string `json:"Redirect"`
}

func fwdQR(w http.ResponseWriter, r *http.Request, data []redirectData) {
	r.ParseForm()

	if len(r.Form) == 0 {
		return
	}

	for _, y := range data {
		if r.Form["name"][0] == y.Name && r.Form["code"][0] == strconv.Itoa(int(y.Code)) {
			http.Redirect(w, r, y.Redirect, 301)
			return
		}
	}

}

func main() {
	var data []redirectData
	dat, err := ioutil.ReadFile("redirects.json")
	if err != nil {
		log.Fatal("Error reading JSON file redirects.json : ", err)
	}

	err = json.Unmarshal(dat, &data)
	if err != nil {
		log.Fatal("Error while Unmarshaling JSON : ", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fwdQR(w, r, data)
	})

	err = http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
