// +build appengine

package main

import (
	"net/http"

	"github.com/ArtnerC/TweetSaver/tsapp"
	"github.com/ArtnerC/TweetSaver/tweetsaver"
	"github.com/codegangsta/martini"

	"appengine"
)

func MapStorage(c martini.Context, r *http.Request) {
	c.MapTo(tsapp.NewDataStore(appengine.NewContext(r)), (*tweetsaver.Persistence)(nil))
}

func ServeStatic() bool {
	return false
}
