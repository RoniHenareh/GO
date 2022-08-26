package main

import (
	"testing"
)

//  go test twopartsum_test.go twopartsum.go

// test that ConcurrentSum sums an even-length array correctly
func TestSumConcurrentCorrectlySumsEvenArray(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := 55

	actual := ConcurrentSum(arr)

	if actual != expected {
		t.Errorf("expected %d, was %d", expected, actual)
	}
}

// test1
func TestSumConcurrentCorrectlySumsEvenArray1(t *testing.T) {
	arr := []int{1, 2, 3, 4, 6, 7, 8, 9, 10}
	expected := 50

	actual := ConcurrentSum(arr)

	if actual != expected {
		t.Errorf("expected %d, was %d", expected, actual)
	}
}

// test2
func TestSumConcurrentCorrectlySumsEvenArray2(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	expected := 66

	actual := ConcurrentSum(arr)

	if actual != expected {
		t.Errorf("expected %d, was %d", expected, actual)
	}
}
