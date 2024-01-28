package web

import (
	"net/http"
	"text/template"

	"github.com/walterdl/gophercises3/story"
)

func serveIndex(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		indexPage(rw, r)
		return
	}

	if r.Method == "POST" {
		chooseArc(rw, r)
		return
	}
}

func indexPage(rw http.ResponseWriter, r *http.Request) {
	arc, err := story.ChooseArc("")
	if err != nil {
		http.Error(rw, "Arc not found", http.StatusBadRequest)
		return
	}

	renderArc(arc, rw, r)
}

func renderArc(arc story.Arc, rw http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFS(staticFolder, "static/index.html")
	if err != nil {
		http.Error(rw, "Error parsing template", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "text/html")
	t.Execute(rw, arc)
}
