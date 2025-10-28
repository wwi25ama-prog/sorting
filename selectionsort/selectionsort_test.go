package selectionsort

import "fmt"

func ExampleSmallestElementIndex() {
	a := []int{5, 1, 8, 3, 4}

	fmt.Println(SmallestElementIndex(a, 0))
	fmt.Println(SmallestElementIndex(a, 2))

	// Output:
	// 1
	// 3
}

func ExampleSelectionSort() {
	a1 := []int{1, 5, 4, 7, 42, 23, 38, -3, 8, 0}
	SelectionSort(a1)
	fmt.Println(a1)

	// Output:
	// [-3 0 1 4 5 7 8 23 38 42]
}
