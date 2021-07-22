package main

func mergeSort(lst []int) {
	size := len(lst)
	if size == 1 {
		lst[0] = lst[0]
	} else if size == 2 {
		if lst[0] > lst[1] {
			lst[0], lst[1] = lst[1], lst[0]
		}
	} else {
		mid := (size - 1) / 2
		firstSlice, finalSlice := lst[0:mid], lst[mid:]
		applySort(firstSlice)
		applySort(finalSlice)
		lst = append(firstSlice, finalSlice...)
	}
}
