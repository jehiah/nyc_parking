package main

import (
	"log"
	"os"

	"github.com/jehiah/nyc_parking/signs"
)

func ParseSigns(filename string, oldFormat bool) []signs.SignPosition {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	var errors []error
	var s []signs.SignPosition
	switch oldFormat {
	case true:
		s, errors = signs.ParseOldCSV(f)
	case false:
		s, errors = signs.ParseCSV(f)
	}

	counter := make(map[string]int)
	seen := make(map[string]bool)
	for _, sign := range s {
		counter[sign.Type.String()] += 1
		counter[sign.Borough] += 1
		if seen[sign.Description] {
			continue
		}
		seen[sign.Description] = true
		// if sign.Type == signs.SpecialInterestParking {
		// 	log.Printf("special: %s", sign.Description)
		// }
		// if sign.Type == signs.UnknownSign {
		// 	log.Printf("unknown: %s", sign.Description)
		// }
		// fmt.Println(sign)
	}
	log.Printf("success:%d failed:%d", len(s), len(errors))
	for _, err := range errors {
		log.Print(err)
	}
	log.Printf("Signs:")
	for t, c := range counter {
		log.Printf("%s: %d", t, c)
	}
	return s
}
