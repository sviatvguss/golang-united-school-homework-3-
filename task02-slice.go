package homework

func reverse(input []int64) (result []int64) {
	out := make([]int64, len(input))
	for i, v := range input {
		out[len(out)-i-1] = v
	}
	result = out
	return
}
