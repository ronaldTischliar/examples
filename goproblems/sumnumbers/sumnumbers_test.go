package sumnumbers_test

import (
	"reflect"
	"testing"

	"github.com/ronaldTischliar/examples/goproblems/sumnumbers"
)

func Test2NumberSum(t *testing.T) {
	t.Run("ok result", func(t *testing.T) {
		expected := []int{11, -1}
		result := sumnumbers.TwoNumberSum([]int{3, 5, -4, 8, 11, 1, -1, 6}, 10)
		if !reflect.DeepEqual(expected, result) {
			t.Errorf("Expected %v got %v", expected, result)
		}
	})

	t.Run("no sum", func(t *testing.T) {
		expected := []int{}
		result := sumnumbers.TwoNumberSum([]int{3, 5, -4, 8, 11, 1, -1, 6}, 20)
		if !reflect.DeepEqual(expected, result) {
			t.Errorf("Expected %v got %v", expected, result)
		}
	})
}

func Test3NumberSum(t *testing.T) {
	t.Run("ok result", func(t *testing.T) {
		expected := [][]int{{-8, 2, 6}, {-8, 3, 5}, {-6, 1, 5}}
		result := sumnumbers.ThreeNumberSum([]int{12, 3, 1, 2, -6, 5, -8, 6}, 0)
		if !reflect.DeepEqual(expected, result) {
			t.Errorf("Expected %v got %v", expected, result)
		}
	})

	t.Run("no sum", func(t *testing.T) {
		expected := [][]int{}
		result := sumnumbers.ThreeNumberSum([]int{12, 3, 1, 2, -6, 5, -8, 6}, 30)
		if !reflect.DeepEqual(expected, result) {
			t.Errorf("Expected %v got %v", expected, result)
		}
	})
}
