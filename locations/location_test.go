package locations

import (
	"fmt"
	"strings"
	"testing"
)

func TestFromCSV(t *testing.T) {
	type testCase struct {
		CSV      string
		Location Location
	}

	tests := []testCase{
		{
			`B,P-200214  ,A J GRIFFIN PLACE,EAST  149 STREET,EAST  144 STREET,E`,
			Location{"B", "P-200214", "A J GRIFFIN PLACE", "EAST 149 STREET", "EAST 144 STREET", "E"},
		},
	}
	for i, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			csv := strings.Split(tc.CSV, ",")
			t.Logf("%#v", csv)
			l := FromCSV(csv)
			if l != tc.Location {
				t.Errorf("got %#v expected %#v", l, tc.Location)
			}
		})
	}
}
