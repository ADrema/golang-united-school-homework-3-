package homework

func average(input [15]float32) (result float32) {
	var sum float32 = 0.0
	var arrayLen = float32(len(input))
	for _, v := range input {
		sum += v
	}
	result = sum / arrayLen
	return result
}
