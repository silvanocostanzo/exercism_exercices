package scrabble

import "strings"

// CalculatePoint takes a letter and returns an integer which represents
// the points for the given letter
func CalculatePoint(l rune) int {
	switch l {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		return 1
	case 'D', 'G':
		return 2
	case 'B', 'C', 'M', 'P':
		return 3
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'K':
		return 5
	case 'J', 'X':
		return 8
	case 'Q', 'Z':
		return 10
	default:
		return 0
	}
}

func Score(word string) int {
	var score int
	formattedWord := strings.ToUpper(word)
	for _, l := range formattedWord {
		score += CalculatePoint(l)
	}
	return score
}
