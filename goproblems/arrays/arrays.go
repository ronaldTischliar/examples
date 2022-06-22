package arrays

import "sort"

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

const HOME_TEAM_WON = 1

func TournamentWinner(competitions [][]string, results []int) string {
	currentBestTeam := ""
	scores := map[string]int{currentBestTeam: 0}
	for idx, competition := range competitions {
		result := results[idx]
		homeTeam, awayTeam := competition[0], competition[1]
		winningTeam := awayTeam
		if result == HOME_TEAM_WON {
			winningTeam = homeTeam
		}
		updateScores(winningTeam, 2, scores)
		if scores[winningTeam] > scores[currentBestTeam] {
			currentBestTeam = winningTeam
		}
	}
	return currentBestTeam
}

func updateScores(team string, points int, scores map[string]int) {
	scores[team] += points
}

func NonConstructibleCoinChange(coins []int) int {
	sort.Ints(coins)
	var currentChangeCreated = 0
	for _, coin := range coins {
		if coin > currentChangeCreated+1 {
			return currentChangeCreated + 1
		}
		currentChangeCreated += coin
	}
	return currentChangeCreated + 1
}
