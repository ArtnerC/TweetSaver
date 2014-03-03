package tweetsave

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

var ErrorBadRequest = errors.New("tweetsave: poorly formatted request")
var ErrorNotFound = errors.New("tweetsave: item not found")

func PerformGet(r http.Request, v View, p Persistence) error {
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		v.DisplayError(ErrorBadRequest, http.StatusBadRequest)
		return ErrorBadRequest
	}

	t := p.Get(id)
	if t == nil {
		v.DisplayError(ErrorNotFound, http.StatusNoContent)
		return ErrorNotFound
	}

	v.DisplayItem(t)
	return nil
}
