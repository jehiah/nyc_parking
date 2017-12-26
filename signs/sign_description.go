package signs

import (
	"strings"
)

var descReplacer = strings.NewReplacer(
	// consistency
	"1 HMP", "1 HOUR METERED PARKING",
	"2 HMP", "2 HOUR METERED PARKING",
	"3 HMP", "3 HOUR METERED PARKING",
	"3HMP", "3 HOUR METERED PARKING",
	"3 HR METERED PARKING", "3 HOUR METERED PARKING",
	"4 HMP", "4 HOUR METERED PARKING",
	"5 HMP", "5 HOUR METERED PARKING",
	"6 HMP", "6 HOUR METERED PARKING",
	"12 HMP", "12 HOUR METERED PARKING",
	"12 NOON", "NOON",
	"NOON 1:3-PM", "NOON TO 1:30PM",
	"8AM 11AM", "8AM TO 11AM",
	"MOON & STAR SYMBOLS", "MOON & STARS SYMBOLS",
	"(HALF MOON&STAR SYMBOLS(", "(HALF MOON&STAR SYMBOLS)",
	"1 HP ", "1 HOUR PARKING ",
	"TRUC LOADING", "TRUCK LOADING",

	// typos
	"NP PARKING", "NO PARKING",
	"HOILDAY", "HOLIDAY",
	"SIGNLE", "SINGLE",
	"SYBBOL", "SYMBOL",
	"BROON", "BROOM",
	"BRROM", "BROOM",
	"SANI-TATION", "SANITATION",
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
	for _, prefix := range []string{
		"NYC PARKING CARD AVAILABLE",
		"PAY AT MUNI-METER",
		"PLAYGROUND",
		"HILL",
		"DIP",
		"SLIPPERY WHEN WET",
		"TWO WAY TRAFFIC",
		"DIVIDED HIGHWAY",
		"PAVEMENT WIDTH",
		"MERGING TRAFFIC",
		"SIGNAL AHEAD",
		"STOP AHEAD",
		"HIDDEN DRIVEWAY",
		"HAZARD",
		"PEDESTRIAN",
		"CHEVRON",
		"REVERSE",
		"DO NOT PASS",
		"WEIGHT",
		"NO TRESPASSING",
		"BICYCLE RIDERS STOP",
		"STAY IN SINGLE LANE",
		"NO TRUCKS",
		"ADVISORY SPEED",
		"GROOVED PAVEMENT",
		"SCHOOL CROSSING",
		"RIGHT LANE ENDS",
		"MERGE SYMBOL",
		"RAILROAD ADVANCE WARNING",
		"DIAMOND DELINEATOR",
		"YELLOW DELINEATOR",
		"DEAD END",
		"HOSPITAL",
		"WINDING ROAD",
		"CURVE SIGN",
		"ROAD NARROWS",
		"LEFT LANE",
		"RIGHT LANE",
		"SIDE ROAD",
		"CROSS ROAD",
		"LARGE ARROW",
		"TURN SIGN",
		"NO SKATEBOARDING",
		"BUS STOP SIGN",
		"PARALLEL PARKING ONLY",
		"ANGLE PARKING ONLY",
		"BACK IN ANGLE PARKING ONLY",
		"HEAD IN ANGLE PARKING ONLY",
		"BIKE (SYMBOL)",
		"ONE WAY",
		"TRUCK ROUTE",
		"THRU TRAFFIC",
		"KEEP",
		"ALL TRAFFIC",
		"NO U-TURN",
		"YIELD",
		"STOP",
		"RESTRICTED LANE",
	} {
		if strings.HasPrefix(d, prefix) {
			return InformationSign
		}
	}

	for _, prefix := range []string{
		"NO ENGINE IDLING",
		"3 MINUTE IDLING",
		"BUS LANE TOW & FINE",
		"BUSES & RIGHT TURNS ONLY",
		"BUSES ONLY",
	} {
		if strings.HasPrefix(d, prefix) {
			return OtherRegulationSign
		}
	}
	
	for _, needle := range[]string{
		"BROOM SYMBOL",
		"(SANITATION SYMBOL",
	} {
		if strings.Contains(d, needle) {
			return StreetCleaning
		}
	}

	// Parking Regulation
	for _, prefix := range []string{
		"NO PARKING",
		"NO STANDING",
		"NO STOPPING",
		"NIGHT REGULATIONS",
		"1 HOUR METERED PARKING",
		"2 HOUR METERED PARKING",
		"3 HOUR METERED PARKING",
		"4 HOUR METERED PARKING",
		"5 HOUR METERED PARKING",
		"6 HOUR METERED PARKING",
		"12 HOUR METERED PARKING",
		"1/2 HOUR PARKING",
		"1 HOUR PARKING",
		"2 HOUR PARKING",
		"3 HOUR PARKING",
		"4 HOUR PARKING",
		"6 HOUR PARKING",
		"48 HOUR PARKING",
		"1 HR MUNI-METER PARKING",
		"2 HR MUNI-METER PARKING",
		"4 HR MUNI-METER PARKING",
		"SPECIAL NIGHT REGULATION",
		"NIGHT REGULATION",
		"TOW AWAY ZONE",
		"TRUCK (SYMBOL) TRUCK LOADING ONLY",
		"BUS (SYMBOL) NON-MTA BUS LAYOVER ONLY",
		"MOON & STARS (SYMBOLS)",
		"PARKING PERMITTED",
		"OTHER TIMES NO PARKING",
	} {
		if strings.HasPrefix(d, prefix) {
			return ParkingSign
		}
	}

	for _, needle := range []string{
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
		"CORPORATION",
		"FORDHAM",
		"NYP",
		"U.S. ",
		"NYCT",
		"ATTORNEY",
		"RELIEF STAND",
		"TAXI / FHV STAND",
		"TAXI STAND",
		"ADMINISTRATION",
		" AVO ",
		" BOARD",
		"CONSUL",
		"DIPLOMAT",
	} {
		if strings.Contains(d, needle) {
			return SpecialInterestParking
		}
	}
	return UnknownSign
}
