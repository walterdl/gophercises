package web

import (
	"embed"
	"fmt"
	"log"
	"net/http"
)

//go:embed static
var staticFolder embed.FS

func Server(port int) {
	http.Handle("/static/", http.FileServer(http.FS(staticFolder)))
	http.HandleFunc("/", serveIndex)

	address := fmt.Sprintf("localhost:%d", port)
	log.Printf("Listening on %s\n", address)
	http.ListenAndServe(address, nil)
}
