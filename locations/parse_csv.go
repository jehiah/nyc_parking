package locations

import (
	"encoding/csv"
	"io"
	"log"
)

func ParseCSV(f io.Reader) ([]Location, []error) {
	r := csv.NewReader(f)
	r.Read() // consume header

	var errors []error
	var locations []Location
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		l, err := FromCSV(row)
		if err != nil {
			errors = append(errors, err)
			continue
		}
		locations = append(locations, l)
	}
	return locations, errors
}
