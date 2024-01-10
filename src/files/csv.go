package readfile

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func OpenCSV(file string) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", rec)
	}
}

// func CreateFile() {
// }
