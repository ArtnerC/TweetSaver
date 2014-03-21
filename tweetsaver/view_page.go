package tweetsaver

import (
	"html/template"
	"net/http"
)

var TemplateDir = "../templates/"

var (
	ItemTemplate     = template.Must(template.ParseFiles(usingMaster(TemplateDir, "Item.html")...))
	ItemListTemplate = template.Must(template.ParseFiles(usingMaster(TemplateDir, "ItemList.html")...))
)

func usingMaster(base, name string) []string {
	return []string{base + "Master.html", base + name}
}

type PageView struct {
	response http.ResponseWriter
}

func NewPageView(w http.ResponseWriter) *PageView {
	return &PageView{response: w}
}

func (pv *PageView) DisplayItem(t *Tweet) {
	if err := ItemTemplate.Execute(pv.response, t); err != nil {
		http.Error(pv.response, err.Error(), http.StatusInternalServerError)
	}
}

func (pv *PageView) DisplayAll(tweets []*Tweet) {
	if err := ItemListTemplate.Execute(pv.response, tweets); err != nil {
		http.Error(pv.response, err.Error(), http.StatusInternalServerError)
	}
}

func (pv *PageView) DisplayError(err error, code int) {
	http.Error(pv.response, err.Error(), code)
}
