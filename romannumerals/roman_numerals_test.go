package romannumerals_test

import (
	"github.com/matryer/is"
	"github.com/mikejeuga/katas/romannumerals"
	"testing"
)

func TestRomanNumeral(t *testing.T) {
	is := is.New(t)
	t.Run("the function should return I when given 1", func(t *testing.T) {

		want := "I"

		got := romannumerals.RomanNumeral(1)
		is.Equal(got, want)
	})

	t.Run("the function should return II when given 2", func(t *testing.T) {

		want := "II"

		got := romannumerals.RomanNumeral(2)
		is.Equal(got, want)
	})

}

