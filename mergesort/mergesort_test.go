package mergesort

import "fmt"

func ExampleMerge() {
	a1 := []int{1, 3, 5, 7}
	a2 := []int{2, 4, 6, 8}

	merged := Merge(a1, a2)
	fmt.Println(merged)

	// Output:
	// [1 2 3 4 5 6 7 8]
}

func ExampleMergeSort() {
	a1 := []int{1, 5, 4, 7, 42, 23, 38, -3, 8, 0}
	a1 = MergeSort(a1)
	fmt.Println(a1)

	// Output:
	// [-3 0 1 4 5 7 8 23 38 42]
}
