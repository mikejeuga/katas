package longestSubstring_test

import (
	"github.com/matryer/is"
	"github.com/mikejeuga/katas/longestSubstring"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	is := is.New(t)
	t.Run("when given abcabcbb it returns 3", func(t *testing.T) {
		given := "abcabcbb"
		expected := 3

		lengthOfLongestSubstring := longestSubstring.LengthOfLongestSubstring(given)
		is.Equal(lengthOfLongestSubstring, expected)
	})

	t.Run("when given bbbbbbb it returns 1", func(t *testing.T) {
		given := "bbbbbbb"
		expected := 1

		lengthOfLongestSubstring := longestSubstring.LengthOfLongestSubstring(given)
		is.Equal(lengthOfLongestSubstring, expected)
	})

	t.Run("when given pwwkew it returns 3", func(t *testing.T) {
		given := "pwwkew"
		expected := 3

		lengthOfLongestSubstring := longestSubstring.LengthOfLongestSubstring(given)
		is.Equal(lengthOfLongestSubstring, expected)
	})


}

