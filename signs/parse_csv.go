package signs

import (
	"encoding/csv"
	"io"
	"log"
)

func ParseCSV(f io.Reader) ([]Sign, []error) {
	r := csv.NewReader(f)
	r.Read() // consume header

	var errors []error
	var signs []Sign
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
