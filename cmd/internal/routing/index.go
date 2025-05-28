package routing

import (
	"net/http"
	"personal/templates"
	"personal/templates/pages"
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	templates.Layout(pages.Index()).Render(r.Context(), w)
}
