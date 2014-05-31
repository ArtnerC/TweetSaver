package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"time"

	ts "github.com/ArtnerC/TweetSaver/tweetsaver"
	"github.com/codegangsta/martini"
)

var m *martini.Martini

func init() {
	flag.Parse()

	//martini.Env = martini.Prod
	m = martini.New()
	m.Use(martini.Recovery())
	m.Use(martini.Logger())
	if ServeStatic() {
		m.Use(martini.Static("static"))
	}
	m.Use(MapStorage)

	r := martini.NewRouter()
	r.Get(`/tweets(\.json)?/:id`, Get)
	r.Get(`/tweets\.html/:id`, GetHTML)

	r.Get(`/tweets(\.json)?`, GetAll)
	r.Get(`/tweets\.html|/index\.html|/`, GetAllHTML)

	r.Get(`/additem\.html`, DisplayAddHTML)
	r.Post(`/tweets.html`, AddHTML)
	r.Post(`/tweets`, Add)

	r.NotFound(NotImplemented)

	m.Action(r.Handle)
	http.Handle("/", m)
}

func Get(params martini.Params, rw http.ResponseWriter, p ts.Persistence) {
	ts.PerformGet(params["id"], ts.NewJSONView(rw), p)
}

func GetHTML(params martini.Params, rw http.ResponseWriter, p ts.Persistence) {
	ts.PerformGet(params["id"], ts.NewPageView(rw), p)
}

func GetAll(rw http.ResponseWriter, req *http.Request, p ts.Persistence) {
	pos, limit := req.FormValue("pos"), req.FormValue("limit")

	if pos == "" {
		ts.PerformGetAll(ts.NewJSONView(rw), p)
	} else {
		ts.PerformGetAt(pos, limit, ts.NewJSONView(rw), p)
	}
}

func GetAllHTML(rw http.ResponseWriter, req *http.Request, p ts.Persistence) {
	pos, limit := req.FormValue("pos"), req.FormValue("limit")

	if pos == "" {
		ts.PerformGetAll(ts.NewPageView(rw), p)
	} else {
		ts.PerformGetAt(pos, limit, ts.NewPageView(rw), p)
	}
}

func DisplayAddHTML(rw http.ResponseWriter) {
	ts.PerformDisplayAdd(ts.NewPageView(rw))
}

func AddHTML(rw http.ResponseWriter, req *http.Request, p ts.Persistence) {
	pv := ts.NewPageViewReq(rw, req)
	err := req.ParseMultipartForm(2048)
	if err != nil {
		log.Printf("AddHTML ParseForm Err: %s", err.Error())
	}

	t := &ts.Tweet{
		Author: req.MultipartForm.Value["Author"][0],
		Text:   req.MultipartForm.Value["Text"][0],
		Link:   req.MultipartForm.Value["Link"][0],
	}
	tweetTime, err := time.ParseInLocation("3:04 PM - 2 Jan 2006", req.MultipartForm.Value["Timestamp"][0], time.Local)
	if err != nil {
		pv.DisplayError(err, http.StatusBadRequest)
		return
	}
	t.Timestamp = tweetTime

	ts.PerformAdd(t, pv, p)
}

func Add(rw http.ResponseWriter, req *http.Request, p ts.Persistence) {
	jv := ts.NewJSONView(rw)
	t := &ts.Tweet{}
	err := json.NewDecoder(req.Body).Decode(t)
	if err != nil {
		jv.DisplayError(err, http.StatusBadRequest)
		return
	}
	ts.PerformAdd(t, jv, p)
}

func NotImplemented(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusNotFound)
	http.ServeFile(rw, req, "static/404.html")
}
