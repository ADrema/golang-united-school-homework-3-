package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	// transfer map into custom type
	inputPair := make(PairList, len(input))
	i := 0
	for k, v := range input {
		inputPair[i] = Pair{k, v}
		i++
	}
	// sort it
	sort.Sort(inputPair)

	for _, v := range inputPair {
		result = append(result, v.Value)
	}
	return result
}

type Pair struct {
	Key   int
	Value string
}

type PairList []Pair

func (inputPair PairList) Len() int {
	return len(inputPair)
}

func (inputPair PairList) Less(i, j int) bool {
	return inputPair[i].Key < inputPair[j].Key
}

func (inputPair PairList) Swap(i, j int) {
	inputPair[i], inputPair[j] = inputPair[j], inputPair[i]
}
