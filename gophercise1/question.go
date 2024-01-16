package main

import (
	"encoding/csv"
	"os"
)

type question struct {
	problem string
	answer  string
}

func readQuestions(f *os.File) ([]question, error) {
	result := make([]question, 0)
	records, err := csvReader(f).ReadAll()

	if err != nil {
		return result, err
	}

	for _, record := range records {
		q := question{
			problem: record[0],
			answer:  record[1],
		}
		result = append(result, q)
	}

	return result, nil
}

func csvReader(f *os.File) *csv.Reader {
	cr := csv.NewReader(f)
	cr.FieldsPerRecord = 2

	return cr
}
