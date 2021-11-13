package quicksort

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

func partionate(a []int) int {
	j := 0
	r := len(a) - 1
	swap(a, 0, r)
	for i := 0; i < len(a); i++ {
		if (a[i] < a[r]) {
			swap(a, j, i)
			j++
		}
	}
	swap(a, j, r)
	return j
}

func QuickSort(a []int) {
	if (len(a) <= 1) {
		return
	}	
	p := partionate(a)
	QuickSort(a[:p])
	QuickSort(a[p+1:])
}