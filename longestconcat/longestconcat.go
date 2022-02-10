package longestconcat

import (
	"strings"
)


func LongestConcat(input []string, k int) string{
	if k > len(input) {
		return ""
	}

	longestString := ""
	concat := Concat(input, k)
	for i := 0; i < len(concat)-k + 1 ; i++ {
		if len(concat[i+1]) > len(concat[i]) {
			longestString = concat[i+1]
		}
	}
	return longestString
}

func Concat(input []string, k int) []string{
	var output []string
	for i := 0; i < len(input)-k + 1 ; i++ {
		combination := strings.Join(input[i:k+i], "")
		output = append(output, combination)
	}
	return output
}