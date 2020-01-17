package main

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {

	if len(src) < 1 {
		return
	}

	temp := make([]int64, len(src))

	// mergeSort(src, 0, len(src)-1, temp)
	mergeSortBottomUp(src, temp)
}

func minInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func mergeSortBottomUp(src, temp []int64) {

	width := 1

	for width = 1; width < len(src); width *= 2 {

		for i := 0; i < len(src); i += 2 * width {
			left, middle, right := i, i+width-1, i+2*width-1

			if right >= len(src) {
				right = minInt(right, len(src)-1)
				middle = minInt(i+width-1, right)
			}

			mergeArray(src, temp, left, middle, right)
		}
	}

}

func mergeSort(src []int64, start, end int, temp []int64) {

	if end == start {
		return
	}

	mid := start + (end-start)/2

	mergeSort(src, start, mid, temp)
	mergeSort(src, mid+1, end, temp)

	mergeArray(src, temp, start, mid, end)
}

func mergeArray(src, temp []int64, start, mid, end int) {
	p1, p2, p3 := start, mid+1, start

	for p1 <= mid && p2 <= end {
		if src[p1] >= src[p2] {
			temp[p3] = src[p2]
			p2++
		} else {
			temp[p3] = src[p1]
			p1++
		}
		p3++
	}

	if p1 <= mid {
		for _, e := range src[p1 : mid+1] {
			temp[p3] = e
			p3++
		}
	} else {
		for _, e := range src[p2 : end+1] {
			temp[p3] = e
			p3++
		}
	}

	copyArray(src, temp, start, end)
}

func copyArray(dest, src []int64, start, end int) {

	for i := start; i <= end; i++ {
		dest[i] = src[i]
	}

}
