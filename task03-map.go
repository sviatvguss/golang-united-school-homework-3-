package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	keys := make([]int, 0, len(input))
	for v := range input {
		keys = append(keys, v)
	}
	sort.Ints(keys)

	result = make([]string, len(input))
	for i, v := range keys {
		result[i] = input[v]
	}
	return
}
