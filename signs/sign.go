package signs

import (
	"fmt"
	"regexp"
	"sort"
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
	Supersedes   []string `json:"supersedes,omitempty"`
	SupersededBy string   `json:"superseded_by,omitempty"`
}

type SignPosition struct {
	Borough  string
	Order    string
	Seq      int
	Distance int
	Arrow    string // "" (both), "W", "E", ...
	Sign
}
type SignPositions []SignPosition

// FilterOrder returns the sign positions for a given Order sorted by Seq
func (s SignPositions) FilterOrder(order string) SignPositions {
	var o SignPositions
	for _, ss := range s {
		if ss.Order == order {
			o = append(o, ss)
		}
	}
	sort.Slice(o, func(i, j int) bool { return o[i].Seq < o[j].Seq })
	return o
}

func (s SignPosition) Summary() string {
	return fmt.Sprintf("Sign[%2d] Distance:%-4d Arrow:%s Sign:%s %s", s.Seq, s.Distance, fmt.Sprintf("%1s", s.Arrow), s.Sign.Type, s.Sign.Description)
}

func FromCSV(row []string) (s SignPosition, err error) {
	// Columns: SRP_Boro,SRP_Order,SRP_seq,SR_Distx,SR_Arrow,Sign_descripition,SR_Mutcd_Code,
	seq, err := strconv.Atoi(row[2])
	if err != nil {
		return s, err
	}
	dist, err := strconv.Atoi(row[3])
	if err != nil {
		return s, err
	}
	s = SignPosition{
		Borough:  row[0],
		Order:    strings.TrimSpace(row[1]),
		Seq:      seq,
		Distance: dist,
		Arrow:    strings.TrimSpace(row[4]),
		Sign: Sign{
			Code: strings.TrimSpace(row[6]),
		},
	}
	if s.Arrow == "NULL" {
		s.Arrow = ""
	}
	description := CleanDescription(strings.TrimSpace(row[5]))
	s.SupersededBy, s.Supersedes = extractSuperseedInfo(description)
	s.Description = deleteSuperseedInfo(description)

	s.Type = SignTypeFromDescription(s.Description)
	return
}

var superseededRegex = regexp.MustCompile(`\((SUPERSEDES|SUPERSEDED BY|USE) ([^\)]+)\)`)

func extractSuperseedInfo(description string) (superseededBy string, superseeds []string) {
	// extract SUPERSEEDED BY, SUPERSEEDS
	for _, matches := range superseededRegex.FindAllStringSubmatch(description, -1) {
		switch matches[1] {
		case "SUPERSEDED BY", "USE":
			fields := strings.Fields(matches[2])
			for _, c := range fields {
				c = strings.Trim(c, ",.")
				if strings.HasPrefix(c, "-") {
					continue
				}
				if strings.Contains(c, "/") {
					// i.e. 5/9/84
					continue
				}
				if strings.Count(c, "-") > 1 {
					// i.e. 11-3-88
					continue
				}
				switch c {
				case "", ",", "&", ".", "DATED", "DON'T":
					continue
				}
				superseededBy = c
				break
			}
		case "SUPERSEDES":
			fields := strings.Fields(matches[2])
			for _, c := range fields {
				c = strings.Trim(c, ",.")
				if strings.HasPrefix(c, "-") {
					continue
				}
				if strings.Contains(c, "/") {
					// i.e. 5/9/84
					continue
				}
				if strings.Count(c, "-") > 1 {
					// i.e. 11-3-88
					continue
				}
				switch c {
				case "", ",", "&", ".", "DATED", "DON'T":
					continue
				}
				superseeds = append(superseeds, c)
			}
		}
	}
	return
}
func deleteSuperseedInfo(description string) string {
	return strings.TrimSpace(superseededRegex.ReplaceAllString(description, " "))
}
