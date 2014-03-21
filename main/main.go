package main

import (
	"github.com/ArtnerC/TweetSaver/simplestore"
	ts "github.com/ArtnerC/TweetSaver/tweetsaver"
	"github.com/codegangsta/martini"
	"net/http"
)

var m *martini.Martini
var Storage = simplestore.NewStorageCache(new(simplestore.FileStorage))

func init() {
	m = martini.New()
	m.Use(martini.Recovery())
	m.Use(martini.Logger())
	m.Use(martini.Static("../static"))

	r := martini.NewRouter()
	r.Get(`/tweets(\.json)?/:id`, Get)
	r.Get(`/tweets\.html/:id`, GetHTML)

	r.Get(`/tweets(\.json)?`, GetAll)
	r.Get(`/tweets\.html|/index\.html|/`, GetAllHTML)
	r.NotFound(NotImplemented)

	m.Action(r.Handle)

	http.Handle("/", m)
}

func main() {
	http.ListenAndServe(":8080", nil)
}

func Get(params martini.Params, rw http.ResponseWriter, req *http.Request) {
	ts.PerformGet(params["id"], ts.NewJSONView(rw), Storage)
}

func GetHTML(params martini.Params, rw http.ResponseWriter, req *http.Request) {
	ts.PerformGet(params["id"], ts.NewPageView(rw), Storage)
}

func GetAll(rw http.ResponseWriter, req *http.Request) {
	ts.PerformGetAll(ts.NewJSONView(rw), Storage)
}

func GetAllHTML(rw http.ResponseWriter, req *http.Request) {
	ts.PerformGetAll(ts.NewPageView(rw), Storage)
}

func NotImplemented(rw http.ResponseWriter, req *http.Request) {
	http.Error(rw, "Not Implemented", http.StatusNotImplemented)
}
