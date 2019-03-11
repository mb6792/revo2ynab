package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type RevoEntry struct {
	Date      string
	Reference string
	PaidOut   string
	Category  string
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatalln("No input file provided")
	}

	input := args[0]

	fmt.Println(input)

	file, _ := os.Open(input)
	reader := csv.NewReader(bufio.NewReader(file))
	reader.LazyQuotes = true
	reader.Comma = ';'

	var entries []RevoEntry
	for {
		line, error := reader.Read()

		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		entries = append(entries, RevoEntry{
			Date:      line[0],
			Reference: line[1],
			PaidOut:   line[2],
			Category:  line[7],
		})
	}
	fmt.Println(entries)

}
