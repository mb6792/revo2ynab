package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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

	entries, error := parse(input)
	if error != nil {
		log.Fatal(error)
	}

	error = write(entries, targetName(input))
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println("Done!")
}

func parse(input string) ([]RevoEntry, error) {
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
			return nil, error
		}

		entries = append(entries, RevoEntry{
			Date:      line[0],
			Reference: line[1],
			PaidOut:   line[2],
			Category:  line[7],
		})
	}

	return entries, nil
}

func write(entries []RevoEntry, name string) error {
	target, _ := os.Create(name)
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
			return error
		}
	}

	return nil
}

func targetName(source string) string {
	return strings.Replace(source, ".csv", "", -1) + "-ynab.csv"
}
