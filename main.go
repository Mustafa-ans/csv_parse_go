package main

import (
	"encoding/csv" // needed to parse csv
	"fmt" // standard input ouput lib
	"log" // lib for logging
	"os"
	"strconv"
	"text/tabwriter"
)

func main() {
	f, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.TabIndent)

	for {
		record, err := r.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatal(err)
		}

		age, err := strconv.Atoi(record[1])
		if err != nil {
			continue
		}

		fmt.Fprintf(w, "%s\t%d\t%s\n", record[0], age, record[2])
	}
	w.Flush()
}
