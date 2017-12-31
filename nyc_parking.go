package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"sort"

	"github.com/jehiah/nyc_parking/signs"
)

type SignCount struct {
	signs.Sign
	Count int `json:"count"`
}

func main() {
	signsFile := flag.String("signs", "", "path to signs.csv")
	locationsFile := flag.String("locations", "", "path to locations.csv")
	signSummary := flag.Bool("sign_summary", false, "output json-lines of sign data")
	flag.Parse()

	s := ParseSigns(*signsFile)

	if *signSummary {
		if len(s) == 0 {
			log.Fatal("no signs")
		}
		signCount := make(map[string]int)
		supersededBy := make(map[string]string)
		signLookup := make(map[string]signs.Sign)

		for _, ss := range s {
			switch ss.Type {
			case signs.BuildingLine, signs.CurbLine:
				continue
			}
			signCount[ss.Code] += 1
			if signCount[ss.Code] > 1 {
				continue
			}
			for _, sss := range ss.Supersedes {
				if supersededBy[sss] == "" {
					supersededBy[sss] = ss.Code
				}
			}
			if ss.SupersededBy != "" && supersededBy[ss.Code] == "" {
				supersededBy[ss.Code] = ss.SupersededBy
			}
			signLookup[ss.Code] = ss.Sign
		}

		log.Printf("%d unique Signs", len(signLookup))
		// collapse counts
		for code := range signLookup {
			target := supersededBy[code]
			if target == "" || target == code {
				continue
			}
			// log.Printf("collapsing %s", sign)
			// recursively find target
			for supersededBy[target] != "" && supersededBy[target] != code && supersededBy[target] != target {
				// log.Printf("recursive %s -> %s", target, supersededBy[target])
				target = supersededBy[target]
			}
			// log.Printf("collapsing %4d %9s -> %9s", signCount[code], code, target)
			signCount[target] += signCount[code]
			delete(signCount, code)
			delete(signLookup, code)
		}
		log.Printf("%d unique Signs (After processing superseeded records)", len(signLookup))

		counts := make([]SignCount, 0, len(signLookup))
		for _, sign := range signLookup {
			counts = append(counts, SignCount{
				Sign:  sign,
				Count: signCount[sign.Code],
			})
		}
		// sort
		sort.Slice(counts, func(i, j int) bool { return counts[i].Count > counts[j].Count })

		for _, sc := range counts {
			b, err := json.Marshal(sc)
			if err != nil {
				log.Fatalf("error %s", err)
			}
			fmt.Printf("%s\n", string(b))
		}

	} else {
		ParseLocations(*locationsFile)
	}
}
