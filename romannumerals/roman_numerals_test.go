package romannumerals_test

import (
	"github.com/matryer/is"
	"github.com/mikejeuga/katas/romannumerals"
	"testing"
)

func TestRomanNumeral(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		description string
		input int
		want string
	}{
		{
			description: "the function should return I when given 1",
			input: 1,
			want: "I",
		},
		{
			description: "the function should return II when given 2",
			input: 2,
			want: "II",
		},
		{
			description: "the function should return III when given 3",
			input: 3,
			want: "III",
		},
		{
			description: "the function should return III when given 4",
			input: 4,
			want: "IV",
		},
	} {
		t.Run(tc.description, func(t *testing.T) {
			is := is.New(t)
			got := romannumerals.RomanNumeral(tc.input)
			is.Equal(got, tc.want)
		})
	}
 }
