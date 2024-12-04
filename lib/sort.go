package lib

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1

	// Choose pivot
	pivot := len(arr) / 2

	// Move pivot to end
	arr[pivot], arr[right] = arr[right], arr[pivot]

	// Subdivide the array
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	// Move pivot to end position
	arr[left], arr[right] = arr[right], arr[left]

	// Recursively sort slices
	QuickSort(arr[:left])
	QuickSort(arr[left+1:])

	return arr
}
