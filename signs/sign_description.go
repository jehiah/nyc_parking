package signs

import (
	"strings"
)

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

	return UnknownSign
}
