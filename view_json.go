package tweetsave

import (
	"encoding/json"
	"io"
	"net/http"
)

type JSONView struct {
	response io.Writer
}

func NewJSONView(w io.Writer) *JSONView {
	return &PageView{response: w}
}

func (jv *JSONView) DisplayItem(t *tweet) {
	err = json.NewEncoder(w).Encode(t)
	if err != nil {
		http.Error(pv.response, err, http.StatusInternalServerError)
		return err
	}
}

func (jv *JSONView) DisplayError(err error, code int) {
	http.Error(pv.response, err, code)
}
