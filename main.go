package main

import (
	"fmt"
	"sync"
)

// binarySearch выполняет бинарный поиск элемента в отсортированном массиве
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1 // Элемент не найден
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := []int{7, 3, 11}

	results := make([]int, len(target))
	var wg sync.WaitGroup

	for i, t := range target {
		wg.Add(1)
		go func(i, t int) {
			defer wg.Done()
			results[i] = binarySearch(arr, t)
		}(i, t)
	}

	wg.Wait()

	for i, t := range target {
		result := results[i]
		if result != -1 {
			fmt.Printf("Элемент %d найден на позиции %d\n", t, result)
		} else {
			fmt.Printf("Элемент %d не найден в массиве\n", t)
		}
	}
}
