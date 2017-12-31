package signs

import (
	"encoding/csv"
	"io"
	"log"
)

func ParseCSV(f io.Reader) ([]SignPosition, []error) {
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
			log.Fatal(err)
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
