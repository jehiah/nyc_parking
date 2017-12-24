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
		// fmt.Println(sign)
	}
	log.Printf("success:%d failed:%d", success, failed)
	for t, c := range counter {
		log.Printf("%s: %d", t, c)
	}
}
