package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type RevoEntry struct {
	Date      string
	Reference string
	PaidOut   string
	PaidIn    string
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
			Date:      strings.TrimSpace(line[0]),
			Reference: line[1],
			PaidOut:   line[2],
			PaidIn:    line[3],
			Category:  line[7],
		})
	}

	return entries, nil
}

func write(entries []RevoEntry, name string) error {
	target, _ := os.Create(name)
	defer target.Close()

	location, _ := time.LoadLocation("UTC")

	writer := csv.NewWriter(target)
	defer writer.Flush()

	writer.Write([]string{
		"Date",
		"Payee",
		"Category",
		"Memo",
		"Outflow",
		"Inflow",
	})

	for _, record := range entries {
		date, _ := time.Parse("Jan 2, 2006", record.Date)
		dateWithYear := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, location)
		dateString := dateWithYear.Format("06/01/02")

		error := writer.Write([]string{
			dateString,
			record.Reference,
			"",
			"",
			record.PaidOut,
			record.PaidIn,
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
