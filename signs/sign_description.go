package signs

import (
	"strings"
)

var descReplacer = strings.NewReplacer(
	"1 HMP", "1 HOUR METERED PARKING",
	"2 HMP", "2 HOUR METERED PARKING",
	"3 HMP", "3 HOUR METERED PARKING",
	"3HMP", "3 HOUR METERED PARKING",
	"3 HR METERED PARKING", "3 HOUR METERED PARKING",
)

func CleanDescription(d string) string {
	return descReplacer.Replace(d)
}

func SignTypeFromDescription(d string) SignType {
	switch d {
	case "":
		return BlankSign
	case "Curb Line":
		return CurbLine
	case "Building Line", "Property Line":
		return BuildingLine
	}
	if strings.HasPrefix(d, "NYC PARKING CARD AVAILABLE") || strings.HasPrefix(d, "NYC PARKING CARD AVAILABLE") {
		return InformationSign
	}
	// NO Par
	for _, prefix := range []string{
		"NO PARKING",
		"NO STANDING",
		"NO STOPPING",
		"NIGHT REGULATIONS",
		"1 HOUR METERED PARKING",
		"2 HOUR METERED PARKING",
		"3 HOUR METERED PARKING",
	} {
		if strings.HasPrefix(d, prefix) {
			return ParkingSign
		}
	}

	for _, kw := range []string{
		"DEP",
		"OFFICE",
		"VEHICLES",
		"JUDGE",
		"FACULTY",
		"POLICE",
		"COURT",
		"COMMISION",
		"FUNUERAL",
		"DOCTOR",
		"JUSTICE",
		"PROSECUTOR",
		"AUTHORITY",
		"CLERK",
		"DHS",
		"STATE",
		"NYS",
		"FEDERAL",
		"COLLEGE",
		"BUREAU",
		"MAYOR",
		"CITY HALL",
		"AMBULANCE",
		"AMBULETTE",
		"BOARD OF",
		"N Y C",
		"MAIL",
	} {
		if strings.Contains(d, kw) {
			return SpecialInterestParking
		}
	}
	return UnknownSign
}
