package urlshort

import (
	"net/http"
	"strings"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.ToLower((*r).URL.Path)

		if value, ok := pathsToUrls[path]; ok {
			w.Header().Set("Location", value)
		}

		w.WriteHeader(301)

		fallback.ServeHTTP(w, r)
	}
}
