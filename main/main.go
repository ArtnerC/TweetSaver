package main

import (
	ts "github.com/ArtnerC/TweetSaver"
	"github.com/codegangsta/martini"
	"net/http"
)

var m *martini.Martini
var Storage = ts.NewStorageCache(new(ts.FileStorage))

func init() {
	m = martini.New()
	m.Use(martini.Recovery())
	m.Use(martini.Logger())
	m.Use(martini.Static("../static"))

	r := martini.NewRouter()
	r.Get(`/get(\.json)?`, Get)
	r.Get(`/get\.html`, GetHTML)
	r.Get(`/getall(\.json)?`, GetAll)
	r.Get(`/getall\.html|/index\.html|/`, GetAllHTML)
	r.NotFound(NotImplemented)

	m.Action(r.Handle)

	http.Handle("/", m)
}

func main() {
	http.ListenAndServe(":8080", nil)
}

func Get(rw http.ResponseWriter, req *http.Request) {
	ts.PerformGet(req, ts.NewJSONView(rw), Storage)
}

func GetHTML(rw http.ResponseWriter, req *http.Request) {
	ts.PerformGet(req, ts.NewPageView(rw), Storage)
}

func GetAll(rw http.ResponseWriter, req *http.Request) {
	ts.PerformGetAll(req, ts.NewJSONView(rw), Storage)
}

func GetAllHTML(rw http.ResponseWriter, req *http.Request) {
	ts.PerformGetAll(req, ts.NewPageView(rw), Storage)
}

func NotImplemented(rw http.ResponseWriter, req *http.Request) {
	http.Error(rw, "Not Implemented", http.StatusNotImplemented)
}
