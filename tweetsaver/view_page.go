package tweetsaver

import (
	"fmt"
	"net/http"

	"github.com/ArtnerC/blueprint"
)

func init() {
	blueprint.MustCompileDir("Master.html", "../templates/")
}

type PageView struct {
	response http.ResponseWriter
	req      *http.Request
}

func NewPageView(w http.ResponseWriter) *PageView {
	return &PageView{response: w}
}

func NewPageViewReq(w http.ResponseWriter, r *http.Request) *PageView {
	return &PageView{response: w, req: r}
}

func (pv *PageView) DisplayItem(t *Tweet) {
	if err := blueprint.Execute(pv.response, "Item.html", t); err != nil {
		http.Error(pv.response, err.Error(), http.StatusInternalServerError)
	}
}

func (pv *PageView) DisplayAll(tweets []*Tweet) {
	if err := blueprint.Execute(pv.response, "ItemList.html", tweets); err != nil {
		http.Error(pv.response, err.Error(), http.StatusInternalServerError)
	}
}

func (pv *PageView) DisplayAddItem() {
	if err := blueprint.Execute(pv.response, "AddItem.html", nil); err != nil {
		http.Error(pv.response, err.Error(), http.StatusInternalServerError)
	}
}

func (pv *PageView) DisplayItemAdded(id int) {
	newURL := fmt.Sprintf("/tweets.html/%d", id)
	http.Redirect(pv.response, pv.req, newURL, http.StatusTemporaryRedirect)
}

func (pv *PageView) DisplayError(err error, code int) {
	http.Error(pv.response, err.Error(), code)
}
