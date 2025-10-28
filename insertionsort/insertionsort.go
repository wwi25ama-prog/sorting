package insertionsort

// MoveLeft bewegt das Element an Stelle i so lange nach links,
// bis es an der richtigen Stelle im bereits sortierten Teil der Liste ist.
func MoveLeft(arr []int, i int) {
	// Schleife: Fange bei i an und tausche das Element mit dem linken Nachbarn,
	// solange es kleiner ist als dieser und solange es nicht am Anfang der Liste ist.
	for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
		arr[j], arr[j-1] = arr[j-1], arr[j]
	}
}

// InsertionSort sortiert die übergebene Liste mittels des Insertion-Sort-Algorithmus.
func InsertionSort(arr []int) {
	// Verwende MoveLeft in einer Schleife.
	for i := 1; i < len(arr); i++ {
		MoveLeft(arr, i)
	}
}
