package main

import (
	"errors"
	"flag"
)

type cliOptions struct {
	filePath string
	limit    int
}

func getOptions() (cliOptions, error) {
	fpath := flag.String("file", "problems.csv", "file to read")
	limit := flag.Int("limit", 30, "time limit for the quiz in seconds")
	flag.Parse()

	result := cliOptions{}
	if *fpath == "" {
		return result, errors.New("file path not specified")
	}

	result.filePath = *fpath
	result.limit = *limit

	return result, nil
}
