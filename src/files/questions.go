package files

import (
	"fmt"
)

func PrintQuestion(questions [][]string) {
	fmt.Println("Write /exit to terminate process")
	for _, line := range questions {
		var input string

		for input != line[1] {
			fmt.Printf("%v   -> ", line[0])
			fmt.Scan(&input)

			if input == "/exit" {
				return
			}
		}
	}
}
