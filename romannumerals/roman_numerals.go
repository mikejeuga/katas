package romannumerals

type RomanToArabic struct {
	Roman string
	Arabic int
}

type ArabicToRoman struct {
	Arabic int
	Roman string
}

var RomanMAP = []RomanToArabic{
	{"M", 1000},
	{"CM", 900},
	{"D", 500},
	{"CD", 400},
	{"C", 100},
	{"XC", 90},
	{"L", 50},
	{"XL", 40},
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

func ArabicNumeral(input string) int {
	var output int
	for RomanNumeral(output) != input {
		output++
	}
		return output
}