package longestconcat

import (
	"strings"
)

func Concat(input []string, k int) []string{
	var output []string
	for i := 0; i < len(input)-k + 1 ; i++ {
		combination := strings.Join(input[i:k+i], "")
		output = append(output, combination)
	}
	return output
}
