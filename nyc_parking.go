package main

import (
	"flag"
)

func main() {
	flag.Parse()
	ParseSigns("data/signs.csv")
	ParseLocations("data/locations.csv")
}
