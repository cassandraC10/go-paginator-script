package main

import "fmt"

func paginate(items []int, page, itemsPerPage int) []int {
	startingIndex := (page - 1) * itemsPerPage
	endIndex := startingIndex + itemsPerPage

	if startingIndex < 0 || startingIndex >= len(items) {
		return []int{}
	}

	if endIndex > len(items) {
		endIndex = len(items)
	}

	return items[startingIndex:endIndex]
}

func main() {
	items:= make([]int, 100)
	for i := 1; i <= 100; i++ {
		items[i-1] = i
	}
	page := 2
	itemsPerPage := 10

	paginatedItems := paginate(items, page, itemsPerPage)
	fmt.Printf("Page %d:\n", page)
	fmt.Println(paginatedItems)
}