package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/walterdl/gophercises3/storage"
	"github.com/walterdl/gophercises3/story"
)

type options struct {
	port  int
	fpath string
}

func main() {
	opts := parseCliOptions()

	err := initStorage(opts)
	if err != nil {
		log.Fatal(err)
	}

	startServer(opts.port)
}

func initStorage(opts options) error {
	s, err := storage.StoryFromFile(opts.fpath)
	if err != nil {
		return err
	}

	story.SetStory(s)
	return nil
}

func startServer(port int) {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", serveIndex)

	address := fmt.Sprintf("localhost:%d", port)
	log.Printf("Listening on %s\n", address)
	http.ListenAndServe(address, nil)
}

func parseCliOptions() options {
	port := flag.Int("port", 8080, "port to listen on")
	fpath := flag.String("file", "story.json", "file to read story from")
	flag.Parse()

	return options{
		port:  *port,
		fpath: *fpath,
	}
}
