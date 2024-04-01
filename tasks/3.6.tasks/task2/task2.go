package task2

func RemoveSpaces(s string) string {
	runes := []rune(s)

	count := 0
	for _, v := range runes {
		if v != ' ' {
			count++
		}
	}

	withoutSpaces := make([]rune, count)

	count = 0
	for _, r := range runes {
		if r != ' ' {
			withoutSpaces[count] = r
			count++
		}
	}

	return string(withoutSpaces)
}
