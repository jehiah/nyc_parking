// Code generated by "stringer -type=SignType"; DO NOT EDIT.

package signs

import "strconv"

const _SignType_name = "UnknownSignCurbLineBuildingLineBusInformationInformationSignParkingSignAngleParkingSpecialInterestParkingBlankSign"

var _SignType_index = [...]uint8{0, 11, 19, 31, 45, 60, 71, 83, 105, 114}

func (i SignType) String() string {
	if i < 0 || i >= SignType(len(_SignType_index)-1) {
		return "SignType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SignType_name[_SignType_index[i]:_SignType_index[i+1]]
}
