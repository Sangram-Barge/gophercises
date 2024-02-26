package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problem struct {
	question string
	answer   string
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "csv file location")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("cannot open file %s \n", *csvFileName))
	}
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		exit(fmt.Sprintf("error parsing file %s \n", *csvFileName))
	}
	problems := parseProblems(lines)
	correctCount := 0
	for i, prob := range problems {
		askQuestion(prob, i, &correctCount)
	}
	fmt.Printf("You scared %d out of %d\n", correctCount, len(problems))
}

func parseProblems(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			question: line[0],
			answer:   line[1],
		}
	}
	return problems
}

func askQuestion(p problem, i int, ansCount *int) {
	fmt.Printf("Problem #1: %d: %s ", i, p.question)
	var ans string
	fmt.Scanf("%s\n", &ans)
	if ans == p.answer {
		*ansCount++
	}
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
