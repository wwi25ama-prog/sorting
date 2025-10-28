package mergesort

import "fmt"

func ExamplePartition() {
	a1 := []int{3, 5, 2, 1, 4}
	a2 := []int{0, 1, 2, 3, 4}
	a3 := []int{4, 3, 2, 1, 0}

	fmt.Println(Partition(a1))
	fmt.Println(Partition(a2))
	fmt.Println(Partition(a3))

	// Output:
	// 2
	// 0
	// 4
}

func ExampleQuickSort() {
	a1 := []int{1, 5, 4, 7, 42, 23, 38, -3, 8, 0}
	a2 := []int{0, 1, 2}
	a3 := []int{2, 1, 0}

	QuickSort(a1)
	QuickSort(a2)
	QuickSort(a3)

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)

	// Output:
	// [-3 0 1 4 5 7 8 23 38 42]
	// [0 1 2]
	// [0 1 2]
}
