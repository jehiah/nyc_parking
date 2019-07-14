package locations

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

type Location struct {
	Borough              string // B, M, ...
	Order                string
	Street               string
	FromStreet, ToStreet string
	Side                 string // N, W, E, S, ...
}
type Locations []Location

var stripWhitespace = regexp.MustCompile(`\s+`)

func FromCSV(row []string) (Location, error) {
	if len(row) != 6 {
		return Location{}, fmt.Errorf("expected %d columns got %d %#v", 6, len(row), row)
	}
	// CSV Columns: boro,order_no,main_st,from_st,to_st,sos
	return Location{
		Borough:    row[0],
		Order:      strings.TrimSpace(row[1]),
		Street:     stripWhitespace.ReplaceAllString(row[2], " "),
		FromStreet: stripWhitespace.ReplaceAllString(row[3], " "),
		ToStreet:   stripWhitespace.ReplaceAllString(row[4], " "),
		Side:       strings.TrimSpace(row[5]),
	}, nil
}

func (l Locations) FindOrder(order string) (Location, bool) {
	for _, ll := range l {
		if ll.Order == order {
			return ll, true
		}
	}
	return Location{}, false
}
func (l Locations) Order(order string) Location {
	if ll, ok := l.FindOrder(order); !ok {
		log.Fatal("order %q not found", order)
	} else {
		return ll
	}
	panic("unreachable")
}
