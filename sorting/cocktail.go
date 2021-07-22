package main

import (
	"fmt"
	"sort"
)

func cocktailSort(lst []int) {
	for i := 0; i < len(lst)-1; i++ {
		if lst[i] > lst[i+1] {
			lst[i], lst[i+1] = lst[i+1], lst[i]
		}
	}
}

func reverseCocktailSort(lst []int) {
	for i := len(lst) - 1; i > 0; i-- {
		if lst[i] < lst[i-1] {
			lst[i], lst[i-1] = lst[i-1], lst[i]
		}
	}
}

func applySort(lst []int) {
	cocktailSort(lst)
	reverseCocktailSort(lst)
}

func main() {
	lst := []int{2, 8, 5, 9, 5, 3, 2, 7, 1}
	fmt.Println("--- Original List ---")
	fmt.Println(lst)

	for i := 1; i < 4; i++ {
		applySort(lst)
	}

	fmt.Println("--- Sorted List ----")
	fmt.Println(lst)
	fmt.Println("Sorted:", sort.IntsAreSorted(lst))
}
