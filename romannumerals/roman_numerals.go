package romannumerals

var RomanMAP =  map[int]string{
	1: "I",
	5: "V",
	10: "X",
	50: "L",
	100: "C",
	500: "D",
	1000: "M",
}


func RomanNumeral(input int) string {
	output := ""

	for k, _ := range RomanMAP {
		if k == input + 1 {
			return RomanMAP[1] + RomanMAP[k]
		}
	}

	for i := 0; i < input; i++ {
		output += "I"
	}


	return output
}