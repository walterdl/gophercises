package main

import (
	"bufio"
	"log"
	"os"
)

type questionResult struct {
	isCorrect bool
	isLastQ   bool
	err       error
}

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
}

func askAll(qa []question, c chan questionResult) {
	for i, q := range qa {
		result := questionResult{}
		isCorrect, err := ask(q)

		if err != nil {
			result.err = err
			c <- result
			return
		}

		result.isCorrect = isCorrect
		result.isLastQ = i == len(qa)-1
		c <- result
	}
}

func ask(q question) (bool, error) {
	log.Println(q.problem, "?")
	scanner.Scan()

	err := scanner.Err()
	if err != nil {
		return false, err
	}

	return scanner.Text() == q.answer, nil
}
