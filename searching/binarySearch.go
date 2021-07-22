package search

import (
	"fmt"
	"sort"
)

func binarySearch(lst []int, target int) int {
	l, r := 0, len(lst)-1
	for l <= r {
		mid := (l + r) / 2

		if target > lst[mid] {
			l = mid + 1
		} else if target < lst[mid] {
			r = mid - 1
		} else {
			return mid
		}
	}

	return -1
}

func main() {
	lst := []int{3, 5, 9, 7, 4, 2, 11}
	sort.Ints(lst)
	fmt.Println("Sorted List:", lst)
	fmt.Println("The index is:", binarySearch(lst, 4))
}
