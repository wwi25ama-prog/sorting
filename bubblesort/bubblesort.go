package bubblesort

// BubbleUp geht einmal durch das Array arr und tauscht benachbarte Elemente,
// wenn das linke größer ist als das rechte.
// Gibt true zurück, wenn mindestens ein Tausch stattgefunden hat.
func BubbleUp(arr []int) bool {
	// Initialisiere eine Variable, um zu verfolgen, ob ein Tausch stattgefunden hat.
	swapped := false
	// Laufe durch das Array und vergleiche benachbarte Elemente.
	// Wenn das linke Element größer ist, tausche sie und setze swapped auf true.
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
			swapped = true
		}
	}
	return swapped
}

// BubbleSort sortiert die übergebene Liste mittels des Bubble-Sort-Algorithmus.
func BubbleSort(arr []int) {
	// Verwende BubbleUp in einer Schleife.
	// Wiederhole den Vorgang, bis keine Vertauschungen mehr stattfinden.
	for BubbleUp(arr) {
	}
}
