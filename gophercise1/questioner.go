package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strings"
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

func askAll(qa []question, shuffle bool, c chan questionResult) {
	if shuffle {
		shuffleQuestions(&qa)
	}

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

func shuffleQuestions(qa *[]question) {
	rand.Shuffle(len(*qa), func(i, j int) {
		(*qa)[i], (*qa)[j] = (*qa)[j], (*qa)[i]
	})
}

func ask(q question) (bool, error) {
	log.Println(q.problem, "?")
	scanner.Scan()

	err := scanner.Err()
	if err != nil {
		return false, err
	}

	return clearAnswer(scanner.Text()) == clearAnswer(q.answer), nil
}

func clearAnswer(a string) string {
	return strings.ToLower(strings.TrimSpace(a))
}
