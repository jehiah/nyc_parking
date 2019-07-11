package main

import (
	"log"
	"os"

	"github.com/jehiah/nyc_parking/locations"
)

func ParseLocations(filename string) []locations.Location {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	l, errors := locations.ParseCSV(f)
	log.Printf("ParseLocations success:%d failed:%d", len(l), len(errors))
	counter := make(map[string]int)
	for _, loc := range l {
		counter[loc.Borough] += 1
	}
	for _, err := range errors {
		log.Print(err)
	}
	log.Printf("Locations by Borough (%d)", len(counter))
	for t, c := range counter {
		log.Printf("%s: %d", t, c)
	}
	return l
}
