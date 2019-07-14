package locations

import (
	"bufio"
	"io"
	"log"
	"strings"
)

func ParseCSV(r io.Reader) (Locations, []error) {
	scanner := bufio.NewScanner(r)
	var errors []error
	var locations []Location
	scanner.Scan() // consume header
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ",")
		l, err := FromCSV(row)
		if err != nil {
			errors = append(errors, err)
			continue
		}
		locations = append(locations, l)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return locations, errors
}
