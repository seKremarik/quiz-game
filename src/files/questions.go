package files

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func ManageArgs() string {
	csvFilePath := flag.String("csv", "src/data/problems.csv", "CSV file in the format 'question,answer'")
	help := flag.Bool("h", false, "Show help message")

	flag.Parse()

	if *help {
		printHelp()
		os.Exit(1)
	}

	absPath, err := filepath.Abs(*csvFilePath)
	if err != nil {
		log.Fatal(err)
	}

	return absPath
}

func PrintQuestions(questions [][]string) {
	fmt.Println("Write /exit to terminate process")
	var input string
	var totalQuestions, totalCorrect int16
	for _, line := range questions {

		// Print the question
		fmt.Print(line[0], " > ")

		_, err := fmt.Scan(&input)
		if err != nil {
			log.Fatal(err)
		}

		// Check answer
		if input == line[1] {
			totalCorrect++
		}

		totalQuestions++
	}

	fmt.Printf("%d questions were asked\n", totalQuestions)
	fmt.Printf("%d were correct answers", totalCorrect)
}

func printHelp() {
	fmt.Printf("Usage of ./quiz:\n-csv string\n\tA CSV file in the format of 'question,answer' (default: 'problems.csv')\n-limit int\n\tThe time limit for the quiz is second (default 10)\n")
}
