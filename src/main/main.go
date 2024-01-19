package main

import (
	quiz "quiz-game/src/files"
)

func main() {
	file := quiz.ManageArgs()
	lines := quiz.OpenCSV(file)

	quiz.PrintQuestions(lines)
}
