package sumnumbers

import "sort"

func TwoNumberSum(array []int, target int) []int {
	nums := map[int]bool{}
	for _, num := range array {
		potentialMatch := target - num
		if _, found := nums[potentialMatch]; found {
			return []int{potentialMatch, num}
		}
		nums[num] = true
	}
	return []int{}
}

func ThreeNumberSum(array []int, target int) [][]int {
	sort.Ints(array)
	result := [][]int{}
	for i := 0; i < len(array)-2; i++ {
		left, right := i+1, len(array)-1
		for left < right {
			currentSum := array[i] + array[left] + array[right]
			if currentSum == target {
				result = append(result, []int{array[i], array[left], array[right]})
				left += 1
				right -= 1
			} else if currentSum < target {
				left += 1
			} else if currentSum > target {
				right -= 1
			}
		}
	}
	return result
}

func FourNumberSum(array []int, target int) [][]int {
	allPairSums := map[int][][]int{}
	result := [][]int{}
	for i := 1; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			currentSum := array[i] + array[j]
			difference := target - currentSum
			if pairs, found := allPairSums[difference]; found {
				for _, pair := range pairs {
					newquad := append(pair, array[i], array[j])
					result = append(result, newquad)
				}
			}
		}
		for k := 0; k < i; k++ {
			currentSum := array[i] + array[k]
			allPairSums[currentSum] = append(allPairSums[currentSum], []int{array[k], array[i]})
		}
	}

	return result
}
