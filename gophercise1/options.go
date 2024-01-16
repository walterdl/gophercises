package main

import (
	"errors"
	"flag"
)

type cliOptions struct {
	filePath string
	limit    int
	shuffle  bool
}

func getOptions() (cliOptions, error) {
	fpath := flag.String("file", "problems.csv", "file to read")
	limit := flag.Int("limit", 30, "time limit for the quiz in seconds")
	shuffle := flag.Bool("shuffle", false, "shuffle the questions")
	flag.Parse()

	result := cliOptions{}
	if *fpath == "" {
		return result, errors.New("file path not specified")
	}

	result.filePath = *fpath
	result.limit = *limit
	result.shuffle = *shuffle

	return result, nil
}
