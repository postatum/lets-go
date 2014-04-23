package views

import (
	"html/template"
	"net/http"
)

func IndexView(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("static/app/index.html")
	context := make(map[string]interface{})

	w.Header().Set("Content-type", "text/html")
	t.Execute(w, &context)
}
