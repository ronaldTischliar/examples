package arrays_test

import (
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
