package singleNumber_test

import (
	"github.com/matryer/is"
	"github.com/mikejeuga/katas/singleNumber"
	"testing"
)

func TestSingleNumber(t *testing.T) {
    t.Parallel()
	for _, tc := range []struct {
		description string
		got []int
		want int
	}{
		{
			description: "should return 1 if given [1]",
			got: []int{1},
			want: 1,
		},
	} {
		t.Run(tc.description, func(t *testing.T) {
			is := is.New(t)
			is.Equal(singleNumber.SingleNumber(tc.got), tc.want)
		})
	}
 }
