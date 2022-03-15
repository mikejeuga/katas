package romannumerals_test

import (
	"github.com/matryer/is"
	"testing"
)

func TestRomanNumeral(t *testing.T) {
	is := is.New(t)
	want := "I"

	got := RomanNumeral(1)
	is.Equal(got, want)

}

func RomanNumeral(i int) string {
	return "I"
}
