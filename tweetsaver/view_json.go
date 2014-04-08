package tweetsaver

import (
	"encoding/json"
	"net/http"
)

type JSONView struct {
	response http.ResponseWriter
}

func NewJSONView(w http.ResponseWriter) *JSONView {
	return &JSONView{response: w}
}

func (jv *JSONView) DisplayItem(t *Tweet) {
	err := json.NewEncoder(jv.response).Encode(t)
	if err != nil {
		http.Error(jv.response, err.Error(), http.StatusInternalServerError)
	}
}

func (jv *JSONView) DisplayAll(tweets []*Tweet) {
	err := json.NewEncoder(jv.response).Encode(tweets)
	if err != nil {
		http.Error(jv.response, err.Error(), http.StatusInternalServerError)
	}
}

func (jv *JSONView) DisplayAddItem() {
	return
}

func (jv *JSONView) DisplayItemAdded(id int) {

}

func (jv *JSONView) DisplayError(err error, code int) {
	http.Error(jv.response, err.Error(), code)
}
