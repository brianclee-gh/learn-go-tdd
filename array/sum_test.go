package main

import (
	"testing"
	"reflect"
)

func TestSum(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, expected int, numbers []int) {
		t.Helper()
		if got != expected {
			t.Errorf("got %d want %d given, %v", got, expected, numbers)
		}
	}

	// t.Run("collection of 5 nums", func(t *testing.T) {
	// 	numbers := []int{1, 2, 3, 4, 5}

	// 	got := Sum(numbers)
	// 	expected := 15
	// 	assertCorrectMessage(t, got, expected, numbers)
	// })

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		expected := 6
		assertCorrectMessage(t, got, expected, numbers)
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		assertCorrectMessage(t, got, want)
	})

	t.Run("handles empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		assertCorrectMessage(t, got, want)
	})

}