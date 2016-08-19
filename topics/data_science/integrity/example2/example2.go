// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example2

// Sample program to compare parsing a clean CSV with Go to
// parsing a clean CSV with python.
package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

func main() {

	// Open the CSV.
	f, err := os.Open("../data/example_clean.csv")
	if err != nil {
		err = errors.Wrap(err, "Could not open CSV")
		log.Fatal(err)
	}

	// Read in the CSV records
	r := csv.NewReader(bufio.NewReader(f))
	records, err := r.ReadAll()
	if err != nil {
		err = errors.Wrap(err, "Could not parse CSV")
		log.Fatal(err)
	}

	// Get the maximum value in the integer column.
	var intMax int
	for _, record := range records {
		intVal, err := strconv.Atoi(record[0])
		if err != nil {
			err = errors.Wrap(err, "Parse failed")
			log.Fatal(err)
		}
		if intVal > intMax {
			intMax = intVal
		}
	}

	// Print the maxium value
	fmt.Println(intMax)
}