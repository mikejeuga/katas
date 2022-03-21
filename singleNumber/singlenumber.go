package singleNumber

func SingleNumber(nums []int) int {
	checkMap := mapBuilder(nums)

	for k, v := range checkMap {
		if v == 1 {
			return k
		}
	}
	return 0
}

func mapBuilder(nums []int) map[int]int {
	checkMap := make(map[int]int)
	for _, num := range nums {
		if checkMap[num] == 0 {
			checkMap[num] = 1
		} else {
			checkMap[num] += 1
		}
	}
	return checkMap
}
