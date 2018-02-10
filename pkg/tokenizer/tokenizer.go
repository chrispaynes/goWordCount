package tokenizer

import (
	"math"
	"strings"
)

// SplitToSlices splits a string into a slice of string slices
func SplitToSlices(str string, w uint) [][]string {
	if len(str) == 0 || w < 1 {
		return [][]string{}
	}

	substrings := strings.Fields(str)

	if w == 1 || w > uint(len(substrings)) {
		return [][]string{
			substrings,
		}
	}

	if uint(len(substrings))%w == 0 {
		return splitEvenly(substrings, uint(w))
	}

	return splitWithRemainder(substrings, uint(w))
}

func splitEvenly(substrings []string, n uint) [][]string {
	var index, low uint
	result := [][]string{}

	for index = 0; index < n; index++ {
		result = append(result, substrings[low:n*(index+1)])
		low += n
	}

	return result
}

func splitWithRemainder(substrings []string, n uint) [][]string {
	var index, low, loops uint
	result := [][]string{}
	loops = uint(math.Ceil(float64(len(substrings)) / float64(n)))

	for index = 0; index < loops; index++ {
		if index == (loops - 1) {
			result = append(result, substrings[low:])
			continue
		}

		result = append(result, substrings[low:n*(index+1)])
		low += n
	}

	return result
}
