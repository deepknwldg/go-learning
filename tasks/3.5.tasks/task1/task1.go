package task1

import (
	"strings"
)

func FrequentWord(str string) string {
	words := strings.Split(str, " ")
	popularWordsStats := make(map[string]int)

	for _, word := range words {
		if _, ok := popularWordsStats[word]; !ok {
			popularWordsStats[word] = 1
		} else {
			popularWordsStats[word]++
		}
	}

	maxCount := 0
	keyResult := ""
	for key, v := range popularWordsStats {
		if maxCount < v {
			maxCount = v
			keyResult = key
		}
	}

	return keyResult
}
