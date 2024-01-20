package urlshort

import (
	"net/http"
	"strings"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.ToLower((*r).URL.Path)

		if url, ok := pathsToUrls[path]; ok {
			redirect(url, w)
			return
		}

		fallback.ServeHTTP(w, r)
	}
}
