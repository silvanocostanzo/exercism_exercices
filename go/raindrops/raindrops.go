package raindrops

import "strconv"

func Convert(number int) string {
	str := ""

	if number%3 == 0 {
		str += "Pling"
	}

	if number%5 == 0 {
		str += "Plang"
	}

	if number%7 == 0 {
		str += "Plong"
	}

	if len(str) == 0 {
		return strconv.Itoa(number)
	}

	return str
}
