package arrays

func IsValidSubsequence(array []int, sequence []int) bool {
	seqIdx := 0
	for _, value := range array {
		if seqIdx == len(sequence) {
			break
		}
		if value == sequence[seqIdx] {
			seqIdx += 1
		}
	}
	return seqIdx == len(sequence)
}

// O(n) time O(n) space
// input array is in ascemding order
func SortedSquaredArray(array []int) []int {
	result := make([]int, len(array))
	smallerValueIdx := 0
	largerValueIdx := len(array) - 1
	for idx := len(array) - 1; idx >= 0; idx-- {
		smallerValue := array[smallerValueIdx]
		largerValue := array[largerValueIdx]
		if abs(smallerValue) > abs(largerValue) {
			result[idx] = smallerValue * smallerValue
			smallerValueIdx += 1
		} else {
			result[idx] = largerValue * largerValue
			largerValueIdx -= 1
		}
	}
	return result
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
