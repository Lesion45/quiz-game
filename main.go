package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	csvFileName := "problems.csv"
	timeLimit := 10

	questions, err := ParseCSV(csvFileName)

	if err != nil {
		log.Fatal("failed to load problem.csv")
	}

	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

	correct := 0

problemLoop:
	for n, value := range questions {
		fmt.Printf("Problem #%d: %s = ", n+1, value[0])
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			break problemLoop
		case answer := <-answerCh:
			if answer == value[1] {
				correct++
			}
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(questions))
}

func ParseCSV(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}
