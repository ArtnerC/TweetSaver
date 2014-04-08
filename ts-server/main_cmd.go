// +build !appengine

package main

import (
	"net/http"

	"github.com/ArtnerC/TweetSaver/simplestore"
	"github.com/ArtnerC/TweetSaver/tweetsaver"
	"github.com/codegangsta/martini"
)

var Storage = simplestore.NewStorageCache(new(simplestore.FileStorage))

func main() {
	http.ListenAndServe(":8080", nil)
}

func MapStorage(c martini.Context, r *http.Request) {
	c.MapTo(Storage, (*tweetsaver.Persistence)(nil))
}

func ServeStatic() bool {
	return true
}
