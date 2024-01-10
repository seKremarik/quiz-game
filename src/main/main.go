package main

import (
	"fmt"
	"log"
	"path/filepath"
	readfile "quiz-game/src/readFile"
)

func main() {
	file, err := filepath.Abs("src/data/problems.csv")
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(file)
	readfile.OpenCSV(file)
}
