package task3

func FrequentRune(str string) rune {
	runes := []rune(str)
	freqRune := make(map[rune]int)

	for _, r := range runes {
		if _, ok := freqRune[r]; ok {
			freqRune[r]++
		} else {
			freqRune[r] = 1
		}
	}

	maxValue := 0
	var resultKey rune
	for key, v := range freqRune {
		if maxValue < v {
			maxValue = v
			resultKey = key
		}
	}

	return resultKey
}
