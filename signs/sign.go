package signs

import (
	"regexp"
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

type Sign struct {
	Type          SignType
	Borough       string
	Order         string
	Seq           int
	Distance      string // TODO convert to int?
	Arrow         string // "" (both), "W", "E", ...
	Description   string
	SignCode      string // Mutcd_Code
	Supersedes    string
	SuperseededBy string
}

var superseededRegex = regexp.MustCompile(`\((SUPERSEDES|SUPERSEDED BY) ([-A-Z0-9]+)\)`)

func FromCSV(row []string) (s Sign, err error) {
	// Columns: SRP_Boro,SRP_Order,SRP_seq,SR_Distx,SR_Arrow,Sign_descripition,SR_Mutcd_Code,
	seq, err := strconv.Atoi(row[2])
	if err != nil {
		return
	}
	s = Sign{
		Borough:     row[0],
		Order:       strings.TrimSpace(row[1]),
		Seq:         seq,
		Distance:    row[3],
		Arrow:       strings.TrimSpace(row[4]),
		Description: CleanDescription(strings.TrimSpace(row[5])),
		SignCode:    strings.TrimSpace(row[6]),
	}
	if s.Arrow == "NULL" {
		s.Arrow = ""
	}
	s.Type = SignTypeFromDescription(s.Description)
	return
}
