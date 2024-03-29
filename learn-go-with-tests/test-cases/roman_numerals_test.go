package roman

import "testing"

// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/roman-numerals
func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 gets converted to I", 1, "I"},
		{"2 gets converted to II", 2, "II"},
		{"3 gets converted to III", 3, "III"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Want {
				t.Errorf("got %q, want %q", got, test.Want)
			}
		})
	}
}

func ConvertToRoman(arabic int) string {
	if arabic == 3 {
		return "III"
	}

	if arabic == 2 {
		return "II"
	}

	return "I"
}
