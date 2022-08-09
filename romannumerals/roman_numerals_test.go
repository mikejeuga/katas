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
			description: "the function should return XICCLXVII when given 1267",
			input: 1267,
			want: "MCCLXVII",
		},
		{
			description: "the function should return MCDVII when given 1407",
			input: 1407,
			want: "MCDVII",
		},

	} {
		t.Run(tc.description, func(t *testing.T) {
			is := is.New(t)
			got := romannumerals.RomanNumeral(tc.input)
			is.Equal(got, tc.want)
		})
	}
 }


func TestEncoder(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		description, input string
		want int
	}{
		{
			description: "the function should return 1 when given I",
			input: "I",
			want: 1,
		},
		{
			description: "the function should return 2 when given II",
			input: "II",
			want: 2,
		},
		{
			description: "the function should return 3 when given III",
			input: "III",
			want: 3,
		},
		{
			description: "the function should return 4 when given IV",
			input: "IV",
			want: 4,
		},
		{
			description: "the function should return 5 when given V",
			input: "V",
			want: 5,
		},
		{
			description: "the function should return 6 when given VI",
			input: "VI",
			want: 6,
		},
		{
			description: "the function should return 122 when given CXXII",
			input: "CXXII",
			want: 122,
		},
		{
			description: "the function should return 150 when given CL",
			input: "CL",
			want: 150,
		},
		{
			description: "the function should return 79 when given LXXIX",
			input: "LXXIX",
			want: 79,
		},
		{
			description: "the function should return 562 when given DLXII",
			input: "DLXII",
			want: 562,
		},
		{
			description: "the function should return 1267 when given XICCLXVII",
			input: "MCCLXVII",
			want: 1267,
		},

	} {
		t.Run(tc.description, func(t *testing.T) {
			is := is.New(t)
			got := romannumerals.ArabicNumeral(tc.input)
			is.Equal(got, tc.want)
		})
	}
}
