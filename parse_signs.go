package main

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"

	"github.com/jehiah/nyc_parking/signs"
)

func main() {
	flag.Parse()

	f, err := os.Open("data/signs.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)
	r.Read() // consume header

	var failed, success int
	counter := make(map[signs.SignType]int)
	seen := make(map[string]bool)
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		sign, err := signs.FromCSV(row)
		if err != nil {
			log.Printf("error %s", err)
			failed += 1
			continue
		}
		success += 1
		counter[sign.Type] += 1
		if seen[sign.Description] {
			continue
		}
		seen[sign.Description] = true
		// if sign.Type == signs.SpecialInterestParking {
		// 	log.Printf("special: %s", sign.Description)
		// }
		if sign.Type == signs.UnknownSign {
			log.Printf("unknown: %s", sign.Description)
		}
		// fmt.Println(sign)
	}
	log.Printf("success:%d failed:%d", success, failed)
	for t, c := range counter {
		log.Printf("%s: %d", t, c)
	}
}
