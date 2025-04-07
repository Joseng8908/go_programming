package quicksort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	input := []int{4, 3, 5, 2, 1}
	want := []int{1, 2, 3, 4, 5}

	got := make([]int, len(input))
	copy(got, input)

	QuickSort(got, 0, len(got)-1)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
