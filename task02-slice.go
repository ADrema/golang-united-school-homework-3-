package homework

func reverse(input []int64) (result []int64) {
	result = make([]int64, len(input))
	var j int = 0
	for i := len(input) - 1; i >= 0; i-- {
		result[j] = input[i]
		j++
	}
	return result
}
