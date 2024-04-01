package task4

import "unicode"

func StringLengthWithoutSpaces(str string) int {
	runes := []rune(str)

	count := 0
	for _, r := range runes {
		if ok := unicode.IsSpace(r); !ok {
			count++
		}
	}

	return count
}
