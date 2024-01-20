package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/walterdl/gophercises/urlshort"
)

type options struct {
	yamlFile string
	dbFile   bool
}

func main() {
	mux := defaultMux()

	opts, err := parseCliOptions()
	if err != nil {
		log.Fatal(err)
	}

	httpHandler, err := resolveHandler(opts, mux)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", httpHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func parseCliOptions() (options, error) {
	yamlFile := flag.String("yaml", "", "YAML file with the URLs to redirect to")
	useDB := flag.Bool("usedb", false, "Use database")
	flag.Parse()

	result := options{}
	if *yamlFile == "" && !*useDB {
		return result, errors.New("specify either -yaml or -usedb")
	}

	result.yamlFile = *yamlFile
	result.dbFile = *useDB

	return result, nil
}

func resolveHandler(opts options, defaultHandler http.Handler) (http.HandlerFunc, error) {
	if opts.dbFile {
		return urlshort.BoltHandler(defaultHandler)
	}

	return urlshort.YAMLHandler(opts.yamlFile, defaultHandler)
}
