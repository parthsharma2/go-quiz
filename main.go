package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	questionsFile := flag.String("csv", "questions.csv", "a csv file in the format 'question, answer'")
	flag.Parse()

	score := 0

	csvFile, err := os.Open(*questionsFile)
	if err != nil {
		errorQuit(fmt.Sprintf("Error occured while opening file %s\n", *questionsFile))
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	lines, err := csvReader.ReadAll()
	if err != nil {
		errorQuit("Error occured while parsing the csv file")
	}

	questions := parseLines(lines)

	for i, question := range questions {
		var ans string
		fmt.Printf("%d) %s: ", i+1, question.ques)
		fmt.Scanf("%s\n", &ans)

		if ans == question.ans {
			score++
		}
	}

	fmt.Printf("You scored: %d/%d\n", score, len(questions))
}

type question struct {
	ques, ans string
}

func parseLines(lines [][]string) []question {
	questions := make([]question, len(lines))

	for i, line := range lines {
		questions[i] = question{
			ques: strings.TrimSpace(line[0]),
			ans:  strings.TrimSpace(line[1]),
		}
	}

	return questions
}

func errorQuit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
