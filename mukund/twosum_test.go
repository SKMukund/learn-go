package main

import (
	"testing"
)

//Testing TwoSum
func TestTwoSum(t *testing.T) {
	result := TwoSumBrute([]int{2, 7, 11, 15}, 9)
	expected := [2]int{0, 1}

	if result != expected {
		t.Fail()
	}
}