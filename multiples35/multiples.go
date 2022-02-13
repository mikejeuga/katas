package multiples35

func Multiples35(i int) int{
	total := 0
	k := i-1
	for k > 0 {
		mod3or5 := k%3 == 0 || k%5 == 0
		if mod3or5 {
			total += k
		}
		k -= 1
	}
	return total
}
