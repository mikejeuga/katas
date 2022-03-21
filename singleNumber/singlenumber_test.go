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
		{
			description: "should return 1 if given [2,2,1]",
			got: []int{2,2,1},
			want: 1,
		},
		{
			description: "should return 4 if given [4,1,2,1,2]",
			got: []int{4,1,2,1,2},
			want: 4,
		},
		{
			description: "should return 5 if given [1,2,5,1,2]",
			got: []int{1,2,5,1,2},
			want: 5,
		},
		{
			description: "should return 7 if given [1,2,7,1,2,3,4,4,3]",
			got: []int{1,2,7,1,2,3,4,4,3},
			want: 7,
		},
	} {
		t.Run(tc.description, func(t *testing.T) {
			is := is.New(t)
			is.Equal(singleNumber.SingleNumber(tc.got), tc.want)
		})
	}
 }
