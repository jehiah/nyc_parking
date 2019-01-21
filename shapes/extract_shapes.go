package main

import (
	"flag"
	"log"
	"fmt"
	"reflect"

	"github.com/jonas-p/go-shp"
)

func main() {
	file := flag.String("filename", "", "")
	flag.Parse()
	
	shape, err := shp.Open(*file)
	if err != nil { 
		log.Fatal(err) 
	} 
	defer shape.Close()
	
	// fields from the attribute table (DBF)
	fields := shape.Fields()
	
	// loop through all features in the shapefile
	var i int
	for shape.Next() {
		i++
		n, p := shape.Shape()
	
		// print feature
		fmt.Println(reflect.TypeOf(p).Elem(), p.BBox())
	
		// print attributes
		for k, f := range fields {
			val := shape.ReadAttribute(n, k)
			fmt.Printf("\t%v: %v\n", f, val)
		}
		fmt.Println()
		if i > 100 {
			break
		}
		
	}

}