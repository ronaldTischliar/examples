package arrays_test

import (
	"reflect"
	"testing"

	"github.com/ronaldTischliar/examples/goproblems/arrays"
)

func TestSequence(t *testing.T) {
	t.Run("sequence", func(t *testing.T) {

		array := []int{5, 1, 22, 3, 7, 25, 6, -1, 8, 10}
		sequence := []int{1, 6, -1, 10}
		isSequence := arrays.IsValidSubsequence(array, sequence)
		if !isSequence {
			t.Errorf("got %v want %v", isSequence, true)
		}

	})

	t.Run("no sequence", func(t *testing.T) {
		array := []int{5, 1, 22, 3, 7, 25, 6, -1, 8, 10}
		sequence := []int{1, 6, -1, -1}
		isSequence := arrays.IsValidSubsequence(array, sequence)
		if isSequence {
			t.Errorf("got %v want %v", isSequence, false)
		}
	})

}

func TestSortedSquaredArray(t *testing.T) {
	t.Run("sequence", func(t *testing.T) {
		input := []int{-6, -3, 1, 2, 5, 8, 9}
		expected := []int{1, 4, 9, 25, 36, 64, 81}
		result := arrays.SortedSquaredArray(input)
		if !reflect.DeepEqual(expected, result) {
			t.Errorf("Expected %v got %v", expected, result)
		}
	})

}

func TestTournamentWinner(t *testing.T) {
	t.Run("winner", func(t *testing.T) {
		competitions := [][]string{
			{"Alfa", "Beta"},
			{"Beta", "Gama"},
			{"Gama", "Alfa"},
		}
		results := []int{1, 0, 0}
		expected := "Alfa"
		winner := arrays.TournamentWinner(competitions, results)
		if winner != expected {
			t.Errorf("got %v want %v", winner, expected)
		}

	})

}

func TestNonConstructibleCoinChange(t *testing.T) {
	t.Run("no coins 20", func(t *testing.T) {
		input := []int{5, 7, 1, 1, 2, 3, 22}
		expected := 20
		result := arrays.NonConstructibleCoinChange(input)
		if result != expected {
			t.Errorf("got %v want %v", result, expected)
		}

	})

}
