package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"flag"
	"io"
	"log"
	"os"
)

func main() {
	err := startQuiz()
	if err != nil {
		log.Fatal(err)
	}
}

func startQuiz() error {
	fpath, err := filePath()
	if err != nil {
		return err
	}
	log.Println("File path:", *fpath)

	file, err := os.Open(*fpath)
	if err != nil {
		return err
	}
	defer file.Close()

	qr := newQuestionReader(file)
	qa := newQuestionAsker()
	quizResult := quizResult{}

	for {
		q, err := qr.read()
		if err != nil {
			if err == io.EOF {
				break
			}

			return err
		}

		answer, err := qa.ask(q.problem)
		if err != nil {
			return err
		}

		quizResult.totalQuestions++
		if answer == q.answer {
			quizResult.successes++
		}
	}

	printResults(quizResult)

	return nil
}

func filePath() (*string, error) {
	fpath := flag.String("file", "problems.csv", "file to read")
	flag.Parse()

	if *fpath == "" {
		return nil, errors.New("file path not specified")
	}

	return fpath, nil
}

type question struct {
	problem string
	answer  string
}

type questionsReader struct {
	csv *csv.Reader
}

func newQuestionReader(f *os.File) *questionsReader {
	csvReader := csv.NewReader(f)
	csvReader.FieldsPerRecord = 2

	return &questionsReader{
		csv: csvReader,
	}
}

func (qr *questionsReader) read() (question, error) {
	q := question{}
	csvRecord, err := (*qr).csv.Read()

	if err != nil {
		return q, err
	}

	q.problem = csvRecord[0]
	q.answer = csvRecord[1]
	return q, nil
}

type questionAsker struct {
	scanner *bufio.Scanner
}

func newQuestionAsker() *questionAsker {
	return &questionAsker{
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func (qa *questionAsker) ask(question string) (string, error) {
	log.Println(question, "?")
	(*qa).scanner.Scan()

	err := (*qa).scanner.Err()
	if err != nil {
		return "", err
	}

	return (*qa).scanner.Text(), nil
}

type quizResult struct {
	totalQuestions int
	successes      int
}

func printResults(r quizResult) {
	log.Println("Total:", r.totalQuestions)
	log.Println("Successes:", r.successes)
	log.Println("Percentage:", float64(r.successes)/float64(r.totalQuestions)*100, "%")
}
