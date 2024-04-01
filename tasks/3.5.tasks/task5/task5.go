// Напишите функцию для преобразования слайса строк в map.
// Ключами в map должны быть строки в слайсе,
// а значениями в map должно быть количество раз,
// которое каждая строка появляется в слайсе.
// Например, для слайса []string{”a”, “b”, “a”} получим map,
// где для ключа “a” будет значение 2, а для ключа “b” будет значение 1.

package task5

func ToFrequencyMap(s []string) map[string]int {
	toMap := make(map[string]int)

	for _, value := range s {
		if _, ok := toMap[value]; !ok {
			toMap[value] = 1
		} else {
			toMap[value]++
		}
	}

	return toMap
}
