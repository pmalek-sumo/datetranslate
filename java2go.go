package datetranslate

import (
	"errors"
	"fmt"
	"strings"
)

// Translate from SimpleDateFormat [1] to ctimefmt [2].
//
// [1] https://docs.oracle.com/javase/8/docs/api/java/text/SimpleDateFormat.html
// [2] https://github.com/observIQ/ctimefmt/blob/4cb1bdfd4b74804fd68c5169d48a8db1e06cfe3e/ctimefmt.go#L63-L97
func SimpleDateFormat2Ctimefmt(in string) (string, error) {
	var sb strings.Builder

	for len(in) > 0 {
		s := getLongestCharacterStreak(in)

		if s[0] == '\'' {
			// Get the string from single quotes.
			out, n, err := getFromQuotes(in)
			if err != nil {
				return "", fmt.Errorf("problem reading from quotes: %w", err)
			}
			if _, err = sb.WriteString(out); err != nil {
				return "", fmt.Errorf("problem writing to string builder: %w", err)
			}
			in = in[n:]

		} else if len(strings.TrimSpace(s)) == 0 {
			// If we've reached the a streak of whitespace just copy it over.
			if _, err := sb.WriteString(s); err != nil {
				return "", fmt.Errorf("problem writing to string builder: %w", err)
			}
			in = in[len(s):]

		} else {
			translated, err := simpleDateFormat2CtimefmtSegment(s)
			if err != nil {
				return "", err
			}
			if _, err = sb.WriteString(translated); err != nil {
				return "", fmt.Errorf("problem writing to string builder: %w", err)
			}
			in = in[len(s):]

		}
	}

	return sb.String(), nil
}

// simpleDateFormat2CtimefmtSegment translates one segment into ctimefmt format.
func simpleDateFormat2CtimefmtSegment(s string) (string, error) {
	switch s {

	// Years
	case "y", "yyy", "yyyy", "Y", "YYY", "YYYY":
		return "%Y", nil
	case "yy", "YY":
		return "%y", nil

	// Months
	case "MMMM":
		return "%B", nil
	case "MMM":
		return "%b", nil
	case "MM":
		return "%m", nil
	case "M":
		return "%q", nil
	case "EEEE":
		return "%A", nil
	case "E", "EE", "EEE":
		return "%a", nil

	// Days
	case "dd":
		return "%d", nil
	case "d":
		return "%g", nil

	// Hours
	case "HH":
		return "%H", nil
	case "KK":
		return "%I", nil
	case "K":
		return "%l", nil

	// Minutes
	case "mm":
		return "%M", nil

	// Seconds
	case "ss":
		return "%S", nil

	// Milliseconds
	case "SSS":
		return "%L", nil

	// AM/PM marker
	case "a":
		return "%p", nil

	// Time Zone
	case "z":
		return "%Z", nil
	case "ZZZZ":
		return "%z", nil
	case "ZZZ":
		return "%z", nil
	case "ZZ":
		return "%z", nil
	case "Z":
		return "%z", nil
	case "XX":
		return "%z", nil

	// Miscellaneous
	case ":", "-", ".", "/", "_", "+":
		return s, nil

	default:
		return "", fmt.Errorf("unknown format string %q", s)
	}
}

// getFromQuotes takes in a string starting with a single quote and extract the
// text enclosed in it, adhering to SimpleDateFormat rules.
func getFromQuotes(in string) (string, int, error) {
	// Firstly trim the string we've received.
	in = strings.TrimPrefix(in, "'")
	end := strings.IndexRune(in, '\'')
	if end == -1 {
		return "", -1, errors.New("invalid format string")
	}
	in = in[:end]

	// Then check if it's just a single character and if so then just return it.
	l := len(in)
	if l <= 1 {
		return in, l + 2, nil // +2 for trimmed quotes
	}

	var (
		sb   strings.Builder
		last = in[0]
	)

	var n int
	if err := sb.WriteByte(in[0]); err != nil {
		return "", -1, err
	}
	n++

	for i := 1; i < l; i++ {
		if in[i] == '\'' && last == '\'' {
			// If we have received 2 single quotes then just append one.
			if err := sb.WriteByte(in[i]); err != nil {
				return "", -1, err
			}

			n += 2
		} else if in[i] != '\'' && last != '\'' {
			// If we received anything else than single quotes than just
			// append it.
			if err := sb.WriteByte(in[i]); err != nil {
				return "", -1, err
			}
			n++
		} else if in[i] == '\'' && i == l-1 {
			// In case we've reached the end of the string and the single quotes
			// were not closed return an error.
			return "", -1, errors.New("invalid format string")
		}

		last = in[i]
	}

	return sb.String(), n + 2, nil // +2 for trimmed quotes
}

// getLongestCharacterStreak returns the longest streak of the same characters
// starting at the beggining of in.
func getLongestCharacterStreak(in string) string {
	l := len(in)

	if l == 0 {
		return ""
	}
	if l == 1 {
		return in
	}

	last := in[0]
	for i := 1; i < l; i++ {
		if in[i] != last {
			return in[:i]
		}
		last = in[i]
	}

	return in
}
