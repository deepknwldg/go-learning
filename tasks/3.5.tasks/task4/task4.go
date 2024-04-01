package task4

func MapKeyIntersect(m1 map[int]struct{}, m2 map[int]struct{}) []int {
	slc := make([]int, 0)

	for key, _ := range m1 {
		if _, ok := m2[key]; ok {
			slc = append(slc, key)
		}
	}

	return slc
}
