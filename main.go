package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatalln("No input file provided")
	}

	input := args[0]

	fmt.Println(input)

	file, _ := os.Open(input)
	reader := csv.NewReader(bufio.NewReader(file))
	reader.Comma = ';'

	for {
		line, error := reader.Read()

		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
	}

}
