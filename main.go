package main

import (
	"github.com/ArtnerC/tweetsave"
	"net/http"
)

func main() {
	http.Handle("/get", Get)

	http.ListenAndServe(":8080", nil)
}

func Get(rw http.ResponseWriter, req *http.Request) {

}
