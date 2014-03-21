package tweetsaver

import (
	"errors"
	"net/http"
	"strconv"
)

var ErrorBadRequest = errors.New("tweetsaver: poorly formatted request")
var ErrorNotFound = errors.New("tweetsaver: item not found")

func PerformGet(idstr string, v View, p Persistence) error {
	id, err := strconv.Atoi(idstr)
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

func PerformGetAll(v View, p Persistence) error {
	tweets := p.GetAll()
	v.DisplayAll(tweets)
	return nil
}