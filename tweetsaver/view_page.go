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
	rw  http.ResponseWriter
	req *http.Request
}

func NewPageView(w http.ResponseWriter) *PageView {
	return &PageView{rw: w}
}

func NewPageViewReq(w http.ResponseWriter, r *http.Request) *PageView {
	return &PageView{rw: w, req: r}
}

func (pv *PageView) DisplayItem(t *Tweet) {
	if err := blueprint.Execute(pv.rw, "Item.html", t); err != nil {
		http.Error(pv.rw, err.Error(), http.StatusInternalServerError)
	}
}

func (pv *PageView) DisplayAll(tweets []*Tweet) {
	if err := blueprint.Execute(pv.rw, "ItemList.html", tweets); err != nil {
		http.Error(pv.rw, err.Error(), http.StatusInternalServerError)
	}
}

func (pv *PageView) DisplayAddItem() {
	if err := blueprint.Execute(pv.rw, "AddItem.html", nil); err != nil {
		http.Error(pv.rw, err.Error(), http.StatusInternalServerError)
	}
}

func (pv *PageView) DisplayItemAdded(id int) {
	newURL := fmt.Sprintf("/tweets.html/%d", id)
	http.Redirect(pv.rw, pv.req, newURL, http.StatusSeeOther)
}

func (pv *PageView) DisplayError(err error, code int) {
	http.Error(pv.rw, err.Error(), code)
}
