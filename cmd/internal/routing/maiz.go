package routing

import (
	"personal/templates"
	"personal/templates/pages"

	"github.com/a-h/templ"
)

func MaizPage() *templ.ComponentHandler {
	return templ.Handler(templates.Layout(pages.Maiz()))
}
