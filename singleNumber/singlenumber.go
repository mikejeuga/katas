package singleNumber

func SingleNumber(input []int) int {
checkMap := make(map[int]int)
	for _, num := range input {
		if checkMap[num] == 0 {
			checkMap[num] = 1
		} else {
			checkMap[num] += 1
		}
	}

	for k, v := range checkMap {
		if v ==1 {
			return k
		}
	}
	return 0
}
