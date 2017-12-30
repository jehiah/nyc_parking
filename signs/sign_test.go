package signs

import (
	"fmt"
	"strings"
	"testing"
)

func TestFromCSV(t *testing.T) {
	type testCase struct {
		CSV  string
		Sign Sign
	}

	tests := []testCase{
		{
			`M,S-263382  ,1,0,NULL,Curb Line,CL        ,`,
			Sign{
				Type:        CurbLine,
				Borough:     "M",
				Order:       "S-263382",
				Seq:         1,
				Distance:    "0",
				Description: "Curb Line",
				SignCode:    "CL",
			},
		},
		{
			`M,S-263382  ,9,556,NULL,NO PARKING (SANITATION BROOM SYMBOL) 11AM TO 12:30PM MON & THURS <----> (SUPERSEDED BY SP-853C),SP-372C   ,`,
			Sign{
				Type:         StreetCleaning,
				Borough:      "M",
				Order:        "S-263382",
				Seq:          9,
				Distance:     "556",
				Description:  "NO PARKING (SANITATION BROOM SYMBOL) 11AM TO 12:30PM MON & THURS <----> (SUPERSEDED BY SP-853C)",
				SignCode:     "SP-372C",
				SupersededBy: "SP-853C",
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
			if s != tc.Sign {
				t.Errorf("got %#v expected %#v", s, tc.Sign)
			}
		})
	}
}
