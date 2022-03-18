package romannumerals

type RomanToArabic struct {
	Roman string
	Arabic int
}

var RomanMAP = []RomanToArabic{
	{"M", 1000},
	{"D", 500},
	{"C", 100},
	{"L", 50},
	{"X", 10},
	{"IX", 9},
	{"V", 5},
	{"IV", 4},
	{"I", 1},
}




func RomanNumeral(input int) string {
	output := ""

	for input > 0 {
		for _, value := range RomanMAP {
			if input >= value.Arabic{
				output += value.Roman
				input -= value.Arabic
				break
			}
		}
	}

	return output
}