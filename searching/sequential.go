package search

import (
	"fmt"
)

func sequentialSearch(lst []int, target int) int {
	for i := 0; i < len(lst); i++ {
		if lst[i] == target {
			return i
		}
	}

	return -1
}

func main() {
	lst := []int{3, 5, 9, 7, 4, 2, 11}
	fmt.Println("The index is:", sequentialSearch(lst, 4))
}
