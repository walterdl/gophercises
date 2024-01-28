package main

import (
	"net/http"

	"github.com/walterdl/gophercises3/story"
)

func chooseArc(rw http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(rw, "Error parsing form", http.StatusBadRequest)
		return
	}

	arcName := r.PostForm.Get("arc")
	if arcName == "" {
		http.Error(rw, "No arc specified", http.StatusBadRequest)
		return
	}

	arc, err := story.ChooseArc(arcName)
	if err != nil {
		http.Error(rw, "Arc not found", http.StatusBadRequest)
		return
	}

	renderArc(arc, rw, r)
}
