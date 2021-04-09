package datetranslate

import (
	"testing"
	"time"

	"github.com/observiq/ctimefmt"
	"github.com/stretchr/testify/require"
)

func TestJava2Go(t *testing.T) {
	testcases := []struct {
		In          string
		Expected    string
		ExpectedErr bool
	}{
		// y
		{
			In:       "yy",
			Expected: "%y",
		},
		{
			In:       "y",
			Expected: "%Y",
		},
		{
			In:       "yyy",
			Expected: "%Y",
		},
		{
			In:       "yyyy",
			Expected: "%Y",
		},
		// Capital Y
		{
			In:       "YY",
			Expected: "%y",
		},
		{
			In:       "Y",
			Expected: "%Y",
		},
		{
			In:       "YYY",
			Expected: "%Y",
		},
		{
			In:       "YYYY",
			Expected: "%Y",
		},
		// Multiple format strings
		{
			In:       "yy yy",
			Expected: "%y %y",
		},
		{
			In:       "y y",
			Expected: "%Y %Y",
		},
		{
			In:       "yyy yyy",
			Expected: "%Y %Y",
		},
		{
			In:       "yyyy yyyy",
			Expected: "%Y %Y",
		},
		{
			In:       "yy  yy",
			Expected: "%y  %y",
		},
		// Month
		{
			In:       "M",
			Expected: "%q",
		},
		{
			In:       "MM",
			Expected: "%m",
		},
		{
			In:       "MMM",
			Expected: "%b",
		},
		{
			In:       "MMMM",
			Expected: "%B",
		},
		// Week: unsupported in ctimefmt.
		{
			In:          "w",
			ExpectedErr: true,
		},
		{
			In:          "ww",
			ExpectedErr: true,
		},
		// Day
		{
			In:       "d",
			Expected: "%g",
		},
		{
			In:       "dd",
			Expected: "%d",
		},
		// Day of the week
		{
			In:       "E",
			Expected: "%a",
		},
		{
			In:       " E",
			Expected: " %a",
		},
		{
			In:       "  E",
			Expected: "  %a",
		},
		{
			In:       "  E ",
			Expected: "  %a ",
		},
		{
			In:       "EE",
			Expected: "%a",
		},
		{
			In:       "EEE",
			Expected: "%a",
		},
		{
			In:       "EEEE",
			Expected: "%A",
		},
		{
			In:       "'Month=' EEEE",
			Expected: "Month= %A",
		},
		{
			In:       "'Month='EEEE",
			Expected: "Month=%A",
		},
		// Hours
		{
			In:       "HH",
			Expected: "%H",
		},
		{
			// Unsupported in ctimefmt.
			In:          "H",
			ExpectedErr: true,
		},
		{
			// Unsupported in ctimefmt.
			In:          "h",
			ExpectedErr: true,
		},
		{
			In:       "K",
			Expected: "%l",
		},
		{
			In:       "KK",
			Expected: "%I",
		},
		// Minutes
		{
			In:       "mm",
			Expected: "%M",
		},
		{
			// Unsupported in ctimefmt.
			In:          "m",
			ExpectedErr: true,
		},
		// Seconds
		{
			In:       "ss",
			Expected: "%S",
		},
		{
			// Unsupported in ctimefmt.
			In:          "s",
			ExpectedErr: true,
		},
		// Milliseconds
		{
			In:       "SSS",
			Expected: "%L",
		},
		// AM/PM marker
		{
			In:       "a",
			Expected: "%p",
		},
		// Time Zone
		{
			In:       "z",
			Expected: "%Z",
		},
		{
			// Unsupported in ctimefmt.
			In:          "zzzz",
			ExpectedErr: true,
		},
		{
			In:       "Z",
			Expected: "%z",
		},
		{
			In:       "ZZ",
			Expected: "%z",
		},
		{
			In:       "ZZZ",
			Expected: "%z",
		},
		{
			In:       "ZZZZ",
			Expected: "%z",
		},
		{
			In:       "XX",
			Expected: "%z",
		},
		{
			// Unsupported in ctimefmt.
			In:          "X",
			ExpectedErr: true,
		},
		{
			// Unsupported in ctimefmt.
			In:          "XXX",
			ExpectedErr: true,
		},
		// Errors
		{
			In:          "Hello",
			ExpectedErr: true,
		},
		{
			In:          "HYY",
			ExpectedErr: true,
		},
		{
			In:          "YYH",
			ExpectedErr: true,
		},
		{
			In:          "'Year yy",
			ExpectedErr: true,
		},
		// Single quotes
		{
			In:       "yyyy 'Year'",
			Expected: "%Y Year",
		},
		{
			In:       "yy 'Year'",
			Expected: "%y Year",
		},
		{
			In:       "'Year' yyyy",
			Expected: "Year %Y",
		},
		{
			In:       "'Year' yy",
			Expected: "Year %y",
		},
		{
			In:       "'Year='yy",
			Expected: "Year=%y",
		},
		// Full examples
		{
			In:       "yyyy-MM-dd HH:mm:ss.SSS",
			Expected: "%Y-%m-%d %H:%M:%S.%L",
		},
		{
			In:       "'Year='yyyy 'Month='MM 'Day='dd",
			Expected: "Year=%Y Month=%m Day=%d",
		},
	}

	now := time.Now()

	for _, tc := range testcases {
		t.Run(tc.In, func(t *testing.T) {
			out, err := SimpleDateFormat2Ctimefmt(tc.In)
			if tc.ExpectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.Expected, out)

				_, err := ctimefmt.Format(tc.Expected, now)
				require.NoError(t, err)
			}
		})
	}

}
