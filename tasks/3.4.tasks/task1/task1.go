package task1

func Reverse(slice []int) []int {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}

	return slice
}

func RemoveDuplicates(slice []int) []int {
	if len(slice) == 0 {
		return nil
	}

	uniqueNumbers := make([]int, 1)
	uniqueNumbers[0] = slice[0]
	for i := 1; i < len(slice); i++ {
		var exist = false
		for j := 0; j < len(uniqueNumbers); j++ {
			if uniqueNumbers[j] == slice[i] {
				exist = true
			}
		}

		if !exist {
			uniqueNumbers = append(uniqueNumbers, slice[i])
		}
		exist = false
	}

	return uniqueNumbers
}
