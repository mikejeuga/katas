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
			description: "the function should return IV when given 4",
			input: 4,
			want: "IV",
		},
		{
			description: "the function should return V when given 5",
			input: 5,
			want: "V",
		},
		{
			description: "the function should return VI when given 6",
			input: 6,
			want: "VI",
		},
		{
			description: "the function should return CXXII when given 122",
			input: 122,
			want: "CXXII",
		},
		{
			description: "the function should return CL when given 150",
			input: 150,
			want: "CL",
		},
		{
			description: "the function should return LXXIX when given 79",
			input: 79,
			want: "LXXIX",
		},
		{
			description: "the function should return DLXII when given 562",
			input: 562,
			want: "DLXII",
		},
		{
			description: "the function should return XICCLXVII when given 11267",
			input: 1267,
			want: "MCCLXVII",
		},

	} {
		t.Run(tc.description, func(t *testing.T) {
			is := is.New(t)
			got := romannumerals.RomanNumeral(tc.input)
			is.Equal(got, tc.want)
		})
	}
 }
