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
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	reader.LazyQuotes = true
	reader.Comma = ';'

	var entries []RevoEntry
	reader.Read()
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

	target, _ := os.Create("ynab.csv")
	defer target.Close()

	writer := csv.NewWriter(target)
	defer writer.Flush()
	for _, record := range entries {
		error := writer.Write([]string{
			record.Date,
			record.Reference,
			record.PaidOut,
		})

		if error != nil {
			log.Fatal(error)
		}
	}
}
