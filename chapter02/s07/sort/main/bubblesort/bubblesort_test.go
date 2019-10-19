package bubblesort

import (
	"fmt"
	"testing"
)

func TestBubbleSort1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	BubbleSort(values)
	fmt.Println(values)
}
func TestBubbleSort2(t *testing.T) {
	values := []int{5}
	BubbleSort(values)
	fmt.Println(values)
}
