package aoc2024

import "slices"

func dayOne(left, right []int) int {
	slices.Sort(left)
	slices.Sort(right)

	distances := make([]int, len(left))
	for i := 0; i < len(left); i++ {
		if left[i] > right[i] {
			distances[i] = left[i] - right[i]
		} else {
			distances[i] = right[i] - left[i]
		}
	}

	result := 0
	for i := 0; i < len(distances); i++ {
		result += distances[i]
	}

	return result
}
