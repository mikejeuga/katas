package romannumerals

type RomanToArabic struct {
	Roman string
	Arabic int
}

var RomanMAP = []RomanToArabic{
	{"X", 10},
	{"IX", 9},
	{"V", 5},
	{"IV", 4},
	{"I", 1},
}




func RomanNumeral(input int) string {
	output := ""

	//numeral, ok := RomanMAP[input]
	//if ok {
	//	return numeral
	//}

	for input > 0 {
		for _, value := range RomanMAP {
			if input >= value.Arabic{
				output += value.Roman
				input -= value.Arabic
			}
		}
	}


	//for i := 0; i < input; i++ {
	//	output += "I"
	//}


	return output
}