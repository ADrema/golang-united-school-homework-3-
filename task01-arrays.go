package homework

func average(input [15]float32) (result float32) {
	result = 0
	for _, v := range input {
		result += v
	}
	return result
}
