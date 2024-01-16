package main

import (
	"log"
	"os"
	"time"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	opts, err := getOptions()
	if err != nil {
		return err
	}

	file, err := os.Open(opts.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	qs, err := readQuestions(file)
	if err != nil {
		return err
	}

	qResultChan := make(chan questionResult)
	timeLimitChan := make(chan bool)
	totalCorrect := 0
	start := time.Now()

	go askAll(qs, qResultChan)
	go wait(opts.limit, timeLimitChan)

Loop:
	for {
		select {
		case qResult := <-qResultChan:
			if qResult.err != nil {
				return qResult.err
			}

			if qResult.isCorrect {
				totalCorrect++
			}

			if qResult.isLastQ {
				break Loop
			}
		case <-timeLimitChan:
			break Loop
		}
	}

	end := time.Now()

	printResults(len(qs), totalCorrect, end.Sub(start))
	return nil
}

func wait(secs int, c chan bool) {
	time.Sleep(time.Duration(secs) * time.Second)
	c <- true
}

func printResults(total int, successes int, timeUsed time.Duration) {
	log.Println("Total:", total)
	log.Println("Successes:", successes)
	log.Println("Percentage:", float64(successes)/float64(total)*100, "%")
	log.Println("Time used (seconds):", timeUsed.Seconds())
}
