package tweetsaver

import (
	"errors"
	"net/http"
	"strconv"
)

var ErrorBadRequest = errors.New("tweetsaver: poorly formatted request")
var ErrorNotFound = errors.New("tweetsaver: item not found")
var ErrorAddFailed = errors.New("tweetsaver: failed adding item")

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

func PerformGetAt(pos, limit string, v View, p Persistence) error {
	intpos, err := strconv.Atoi(pos)
	if intpos < 0 || err != nil {
		v.DisplayError(ErrorBadRequest, http.StatusBadRequest)
		return ErrorBadRequest
	}

	intlimit, err := strconv.Atoi(limit)
	if intlimit <= 0 || err != nil {
		intlimit = 10
	}

	tweets := p.GetAt(intpos, intlimit)
	v.DisplayAll(tweets)
	return nil
}

func PerformGetAll(v View, p Persistence) error {
	tweets := p.GetAll()
	v.DisplayAll(tweets)
	return nil
}

func PerformDisplayAdd(v View) {
	v.DisplayAddItem()
}

func PerformAdd(t *Tweet, v View, p Persistence) error {
	id, err := p.Add(t)
	if err != nil {
		v.DisplayError(ErrorAddFailed, http.StatusInternalServerError)
		return ErrorAddFailed
	}

	v.DisplayItemAdded(id)
	return nil
}
