package qsort

import (
	"fmt"
	"testing"
)

func TestQuickSort1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	QuickSort(values)
	fmt.Println(values)
}
func TestQuickSort2(t *testing.T) {
	values := []int{5}
	QuickSort(values)
	fmt.Println(values)
}
