package romannumerals

func RomanNumeral(input int) string {
	output := ""

	if input == 4 {
		output = "IV"
		return output
	}
	for i := 0; i < input; i++ {
		output += "I"
	}


	return output
}
