package bubblesort

import "fmt"

func ExampleBubbleUp() {
	a := []int{5, 1, 8, 3, 4}

	fmt.Println(BubbleUp(a))
	fmt.Println(a)

	fmt.Println(BubbleUp(a))
	fmt.Println(a)

	fmt.Println(BubbleUp(a))
	fmt.Println(a)

	// Output:
	// true
	// [1 5 3 4 8]
	// true
	// [1 3 4 5 8]
	// false
	// [1 3 4 5 8]
}

func ExampleBubbleSort() {
	a1 := []int{1, 5, 4, 7, 42, 23, 38, -3, 8, 0}
	BubbleSort(a1)
	fmt.Println(a1)

	// Output:
	// [-3 0 1 4 5 7 8 23 38 42]
}
