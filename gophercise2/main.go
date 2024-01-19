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
}

func main() {
	mux := defaultMux()

	opts, err := parseCliOptions()
	if err != nil {
		log.Fatal(err)
	}

	mapHandler, err := urlshort.YAMLHandler(opts.yamlFile, mux)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mapHandler)
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
	flag.Parse()

	result := options{}
	if *yamlFile == "" {
		return result, errors.New("please provide a YAML file with the URLs to redirect to")
	}

	result.yamlFile = *yamlFile

	return result, nil
}
