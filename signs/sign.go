package signs

import (
	"strconv"
	"strings"
)

//go:generate stringer -type=SignType
type SignType int8

const (
	UnknownSign SignType = iota
	CurbLine
	BuildingLine
	BusInformation
	InformationSign
	ParkingSign
	AngleParking
	SpecialInterestParking
	BlankSign
	OtherRegulationSign
	StreetCleaning
)

type RawSign struct {
	Type        SignType
	Borough     string
	Order       string
	Seq         int
	Distance    string // TODO
	Arrow       string
	Description string
	Mutcd_Code  string
}

func FromCSV(row []string) (s RawSign, err error) {
	seq, err := strconv.Atoi(row[2])
	if err != nil {
		return
	}
	s = RawSign{
		Borough:     row[0],
		Order:       strings.TrimSpace(row[1]),
		Seq:         seq,
		Distance:    row[3],
		Arrow:       row[4],
		Description: CleanDescription(strings.TrimSpace(row[5])),
		Mutcd_Code:  strings.TrimSpace(row[6]),
	}
	s.Type = SignTypeFromDescription(s.Description)
	return
}
