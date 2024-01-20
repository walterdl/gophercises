package urlshort

import "net/http"

func redirect(url string, w http.ResponseWriter) {
	w.Header().Set("Location", url)
	w.WriteHeader(301)
}
