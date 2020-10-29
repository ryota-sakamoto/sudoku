package testutil

func CreateSequence(list ...int) [9]map[int]struct{} {
	result := [9]map[int]struct{}{}

	for k, v := range list {
		if result[k] == nil {
			result[k] = map[int]struct{}{}
		}

		result[k][v] = struct{}{}
	}

	return result
}
