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

func (s SignType) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}

type Sign struct {
	Type         SignType `json:"type"`
	Description  string   `json:"description"`
	Code         string   `json:"code"` // Mutcd_Code
	Supersedes   string   `json:"supersedes,omitempty"`
	SupersededBy string   `json:"superseded_by,omitempty"`
}

type SignPosition struct {
	Borough  string
	Order    string
	Seq      int
	Distance string // TODO convert to int?
	Arrow    string // "" (both), "W", "E", ...
	Sign
}

var superseededRegex = regexp.MustCompile(`\((SUPERSEDES|SUPERSEDED BY|USE) ([-A-Z0-9]+)\)`)

func FromCSV(row []string) (s SignPosition, err error) {
	// Columns: SRP_Boro,SRP_Order,SRP_seq,SR_Distx,SR_Arrow,Sign_descripition,SR_Mutcd_Code,
	seq, err := strconv.Atoi(row[2])
	if err != nil {
		return
	}
	s = SignPosition{
		Borough:  row[0],
		Order:    strings.TrimSpace(row[1]),
		Seq:      seq,
		Distance: row[3],
		Arrow:    strings.TrimSpace(row[4]),
		Sign: Sign{
			Code: strings.TrimSpace(row[6]),
		},
	}
	if s.Arrow == "NULL" {
		s.Arrow = ""
	}
	description := CleanDescription(strings.TrimSpace(row[5]))

	// extract SUPERSEEDED BY, SUPERSEEDS
	for _, matches := range superseededRegex.FindAllStringSubmatch(description, -1) {
		switch matches[1] {
		case "SUPERSEDED BY", "USE":
			s.SupersededBy = matches[2]
		case "SUPERSEDES":
			s.Supersedes = matches[2]
		}
	}
	s.Description = strings.TrimSpace(superseededRegex.ReplaceAllString(description, " "))

	s.Type = SignTypeFromDescription(s.Description)
	return
}
