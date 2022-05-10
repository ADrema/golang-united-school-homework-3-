package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	sort.Sort(SortedByKeys(input))
	for _, v := range input {
		result = append(result, v)
	}
	return result
}

type SortedByKeys map[int]string

func (input SortedByKeys) Len() int {
	return len(input)
}

func (input SortedByKeys) Less(i, j int) bool {
	return i < j
}

func (input SortedByKeys) Swap(i, j int) {
	input[i], input[j] = input[j], input[i]
}
