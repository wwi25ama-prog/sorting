package selectionsort

// SmallestElementIndex gibt den Index des kleinsten Elements im Array arr ab Index start zurück.
func SmallestElementIndex(arr []int, start int) int {
	minIndex := start
	for i := start + 1; i < len(arr); i++ {
		if arr[i] < arr[minIndex] {
			minIndex = i
		}
	}
	return minIndex
}

// SelectionSort sortiert die übergebene Liste mittels des Selection-Sort-Algorithmus.
func SelectionSort(arr []int) {
	// Verwende SmallestElementIndex in einer Schleife.
	// Tausche das aktuelle Element mit dem kleinsten gefundenen Element.
	for i := 0; i < len(arr)-1; i++ {
		minIndex := SmallestElementIndex(arr, i)
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
