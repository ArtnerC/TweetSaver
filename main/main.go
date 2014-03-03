package main

import (
	ts "github.com/ArtnerC/TweetSaver"
	"net/http"
)

var Storage = ts.NewStorageCache(new(ts.FileStorage))

func main() {
	http.HandleFunc("/get", Get)
	http.HandleFunc("/get.json", Get)
	http.HandleFunc("/get.html", GetHTML)

	http.ListenAndServe(":8080", nil)
}

func Get(rw http.ResponseWriter, req *http.Request) {
	ts.PerformGet(req, ts.NewJSONView(rw), Storage)
}

func GetHTML(rw http.ResponseWriter, req *http.Request) {
	ts.PerformGet(req, ts.NewPageView(rw), Storage)
}
