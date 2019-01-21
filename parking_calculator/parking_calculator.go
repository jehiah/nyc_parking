package main

import (
	"context"
	"log"
	"os"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/jehiah/nyc_parking/internal/coord/client"
	"github.com/jehiah/nyc_parking/internal/coord/client/operations"
)

//
func authFunc(c runtime.ClientRequest, r strfmt.Registry) error {
	c.SetQueryParam("access_key", os.Getenv("COORD_TOKEN"))
	return nil
}

var auth = runtime.ClientAuthInfoWriterFunc(authFunc)

func main() {
	// https://coord.co/docs/searchcurbs#/introduction/getting-all-rules/sample-response
	ctx := context.Background()

	c := client.Default
	d, err := c.Operations.GetAtTimeByBounds(&operations.GetAtTimeByBoundsParams{
		MaxLatitude:  40.915568,
		MinLatitude:  40.90, // 40.495992
		MaxLongitude: -74.257159,
		MinLongitude: -73.699215,
		Context:      ctx,
	}, auth)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(len(d.Payload.Features))
}
