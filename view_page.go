package tweetsave

import (
	"html/template"
	"io"
	"net/http"
)

type PageView struct {
	response io.Writer
}

func NewPageView(w io.Writer) *PageView {
	return &PageView{response: w}
}

func (pv *PageView) DisplayItem(t *tweet) {
	if err := ItemTemplate.Execute(pv.response, t); err != nil {
		http.Error(pv.response, err, http.StatusInternalServerError)
	}
}

func (pv *PageView) DisplayError(err error, code int) {
	http.Error(pv.response, err, code)
}

var ItemHTML = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<title>Tweet {{.Id}}</title>
	</head>
	<body>
	<h1>{{.Author}}</h1>
	<p>{{.Text}}</p>
	<a href="{{.Link}}">View on Twitter</a>
	</body>
</html>
`
var ItemTemplate = template.Must(template.New("ItemHTML").Parse(ItemHTML))
