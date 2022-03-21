package singleNumber

func SingleNumber(got []int) int {
	for _, num := range got {
		if num == 1 {
			return num
		}
	}
	return got[0]
}
