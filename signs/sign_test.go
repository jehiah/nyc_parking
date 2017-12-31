package signs

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestFromCSV(t *testing.T) {
	type testCase struct {
		CSV  string
		Sign SignPosition
	}

	tests := []testCase{
		{
			`M,S-263382  ,1,0,NULL,Curb Line,CL        ,`,
			SignPosition{
				Borough:  "M",
				Order:    "S-263382",
				Seq:      1,
				Distance: "0",
				Sign: Sign{
					Type:        CurbLine,
					Description: "Curb Line",
					Code:        "CL",
				},
			},
		},
		{
			`M,S-263382  ,9,556,NULL,NO PARKING (SANITATION BROOM SYMBOL) 11AM TO 12:30PM MON & THURS <----> (SUPERSEDED BY SP-853C),SP-372C   ,`,
			SignPosition{
				Borough:  "M",
				Order:    "S-263382",
				Seq:      9,
				Distance: "556",
				Sign: Sign{
					Type:         StreetCleaning,
					Description:  "NO PARKING (SANITATION BROOM SYMBOL) 11AM TO 12:30PM MON & THURS <---->",
					Code:         "SP-372C",
					SupersededBy: "SP-853C",
				},
			},
		},
	}
	for i, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			csv := strings.Split(tc.CSV, ",")
			t.Logf("%#v", csv)
			s, err := FromCSV(csv)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(s, tc.Sign) {
				t.Errorf("got %#v expected %#v", s, tc.Sign)
			}
		})
	}
}

func TestExtractSuperseedInfo(t *testing.T) {
	type testCase struct {
		Description  string
		SupersededBy string
		Superseeds   []string
	}

	tests := []testCase{
		{
			`(SUPERSEDED BY SP-464C DATED 8-2-99)`,
			"SP-464C",
			nil,
		},
	}

	for i, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Log(tc.Description)
			sb, ss := extractSuperseedInfo(tc.Description)
			if sb != tc.SupersededBy {
				t.Errorf("got SupersededBy %q expected %q", sb, tc.SupersededBy)
			}
			if !reflect.DeepEqual(ss, tc.Superseeds) {
				t.Errorf("got Superseeds %q expected %q", ss, tc.Superseeds)
			}
		})
	}
}
