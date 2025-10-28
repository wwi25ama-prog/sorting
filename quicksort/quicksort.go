package mergesort

// Partition teilt das Array arr in zwei Teile, basierend auf dem ersten Element
// (dem Pivot). Alle Elemente kleiner oder gleich dem Pivot werden nach links,
// alle größeren nach rechts einsortiert. Gibt die neue Position des Pivotelements zurück.
func Partition(arr []int) int {
	// Verwende zwei Indizes für die linke und rechte Position im Array.
	// Lasse diese Indizes so lange aufeinander zugehen, bis sie sich treffen.
	// Tausche Elemente, die auf der falschen Seite des Pivots stehen.
	i, j := 0, len(arr)-1
	pivot := arr[0]

	for i < j {
		for i < j && arr[j] > pivot {
			j--
		}
		for i < j && arr[i] < pivot {
			i++
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	return i
}

// QuickSort sortiert die übergebene Liste mittels des Quick-Sort-Algorithmus.
func QuickSort(arr []int) {
	// Basisfall: Wenn die Liste weniger als 2 Elemente hat, ist sie bereits sortiert.
	if len(arr) < 2 {
		return
	}

	// Teile das Array mittels Partition und sortiere die beiden Teile rekursiv.
	pivotIndex := Partition(arr)
	QuickSort(arr[:pivotIndex])
	QuickSort(arr[pivotIndex+1:])
}
