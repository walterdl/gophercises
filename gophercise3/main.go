package main

import (
	"flag"
	"log"

	"github.com/walterdl/gophercises3/storage"
	"github.com/walterdl/gophercises3/story"
	"github.com/walterdl/gophercises3/web"
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

	web.Server(opts.port)
}

func initStorage(opts options) error {
	s, err := storage.StoryFromFile(opts.fpath)
	if err != nil {
		return err
	}

	story.SetStory(s)
	return nil
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
