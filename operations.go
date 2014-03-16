package tweetsaver

import (
	"errors"
	"net/http"
	"strconv"
)

var ErrorBadRequest = errors.New("tweetsaver: poorly formatted request")
var ErrorNotFound = errors.New("tweetsaver: item not found")

func PerformGet(r *http.Request, v View, p Persistence) error {
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		v.DisplayError(ErrorBadRequest, http.StatusBadRequest)
		return ErrorBadRequest
	}

	t := p.Get(id)
	if t == nil {
		v.DisplayError(ErrorNotFound, http.StatusNotFound)
		return ErrorNotFound
	}

	v.DisplayItem(t)
	return nil
}

func PerformGetAll(req *http.Request, v View, p Persistence) error {
	tweets := p.GetAll()
	v.DisplayAll(tweets)
	return nil
}
