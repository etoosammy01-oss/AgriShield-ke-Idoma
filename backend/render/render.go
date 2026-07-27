package render

import (
	"html/template"
	"net/http"
)

var TMPL = template.Must(
	template.ParseGlob("template/*.html"),
)

func RenderTemplates(w http.ResponseWriter, name string, data any) error {
	return TMPL.ExecuteTemplate(w, name, data)
}
