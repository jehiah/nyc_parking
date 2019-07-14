package signs

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"strings"
)

func ParseCSV(f io.Reader) (SignPositions, []error) {
	r := csv.NewReader(f)
	r.Read() // consume header

	var errors []error
	var signs []SignPosition
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("parseCSV %s", err)
		}

		sign, err := FromCSV(row)
		if err != nil {
			errors = append(errors, err)
			continue
		}
		signs = append(signs, sign)
	}
	return signs, errors
}

// ParseOldCSV handles CSV's didn't format the description properly,
// so strip the last field, then the first and anything
// remaining is the description
func ParseOldCSV(f io.Reader) (SignPositions, []error) {
	r := bufio.NewScanner(f)
	r.Scan() // consume header
	var errors []error
	var signs []SignPosition
	for r.Scan() {
		var row []string
		line := r.Text()
		chunks := strings.Split(line, ",")
		row = chunks[:5]
		row = append(row, strings.Join(chunks[5:len(chunks)-1], ","))
		row = append(row, chunks[len(chunks)-1])
		sign, err := FromCSV(row)
		if err != nil {
			errors = append(errors, err)
			continue
		}
		signs = append(signs, sign)
	}
	if err := r.Err(); err != nil {
		log.Fatalf("parseOldCSV %s", err)
	}
	return signs, errors
}
