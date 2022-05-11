package homework

func average(input [15]float32) (result float32) {
	N := len(input)
	var sum float32

	for _, v := range input {
		sum += v
	}
	result = sum / float32(N)
	return
}
