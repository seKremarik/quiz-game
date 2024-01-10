package main

import (
	"log"
	"path/filepath"
	quiz "quiz-game/src/files"
)

func main() {
	file, err := filepath.Abs("src/data/problems.csv")
	if err != nil {
		log.Panic(err)
	}

	quiz.OpenCSV(file)
}
