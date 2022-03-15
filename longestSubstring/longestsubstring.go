package longestSubstring

import (
	"strings"
)

func LengthOfLongestSubstring(s string) int {
	substring := ""
	split := strings.Split(s, "")

	for i := 0; i < len(s)-1; i++ {
		if !strings.Contains(substring, split[i]) && split[i] != split[i+1] {
			substring += split[i]
		}
	}

	return len(substring)
}


