package files

import (
	"encoding/csv"
	"log"
	"os"
)

func OpenCSV(file string) [][]string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	rec, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return rec
}
