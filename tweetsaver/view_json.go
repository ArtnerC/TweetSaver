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
		jv.DisplayError(err, http.StatusInternalServerError)
	}
}

func (jv *JSONView) DisplayResults(tweets []*Tweet, pos, total int) {
	s := struct {
		Tweets []*Tweet
		Length int
		Pos    int
		Total  int
	}{
		tweets,
		len(tweets),
		pos,
		total,
	}
	err := json.NewEncoder(jv.response).Encode(&s)
	if err != nil {
		jv.DisplayError(err, http.StatusInternalServerError)
	}
}

func (jv *JSONView) DisplayAll(tweets []*Tweet) {
	err := json.NewEncoder(jv.response).Encode(tweets)
	if err != nil {
		jv.DisplayError(err, http.StatusInternalServerError)
	}
}

func (jv *JSONView) DisplayAddItem() {
	return
}

func (jv *JSONView) DisplayItemAdded(id int) {
	jsonID := struct{ id int }{id: id}
	if err := json.NewEncoder(jv.response).Encode(&jsonID); err != nil {
		jv.DisplayError(err, http.StatusInternalServerError)
	}
}

func (jv *JSONView) DisplayError(err error, code int) {
	e := struct {
		e string `json:"error"`
	}{e: err.Error()}
	ebyte, _ := json.Marshal(&e)
	//e := fmt.Sprintf(`{ "error": "%s" }`, err.Error())
	http.Error(jv.response, string(ebyte), code)
}
