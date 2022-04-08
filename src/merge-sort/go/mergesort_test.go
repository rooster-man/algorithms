package mergesort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	actual := MergeSort([]int{10, 7, 9, 6, 8, 4, 5, 2, 3, 1})

	if reflect.DeepEqual(expected, actual) != true {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}
