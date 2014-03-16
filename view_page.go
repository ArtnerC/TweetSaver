package tweetsaver

import (
	"html/template"
	"net/http"
)

type PageView struct {
	response http.ResponseWriter
}

func NewPageView(w http.ResponseWriter) *PageView {
	return &PageView{response: w}
}

func (pv *PageView) DisplayItem(t *tweet) {
	if err := ItemTemplate.Execute(pv.response, t); err != nil {
		http.Error(pv.response, err.Error(), http.StatusInternalServerError)
	}
}

func (pv *PageView) DisplayAll(tweets []*tweet) {
	if err := ItemListTemplate.Execute(pv.response, tweets); err != nil {
		http.Error(pv.response, err.Error(), http.StatusInternalServerError)
	}
}

func (pv *PageView) DisplayError(err error, code int) {
	http.Error(pv.response, err.Error(), code)
}

var ItemTemplate = template.Must(template.New("Item.html").ParseFiles("../html/Item.html"))
var ItemListTemplate = template.Must(template.New("ItemList.html").ParseFiles("../html/ItemList.html"))
