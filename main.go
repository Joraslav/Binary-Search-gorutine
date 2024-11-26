package main

import (
	"fmt"
	"sync"
)

type Result struct {
	Index int
	Value int
}

func binarySearch(arr []int, target int, index int, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			results <- Result{Index: index, Value: mid}
			return
		}

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	results <- Result{Index: index, Value: -1}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := []int{7, 3, 11}

	results := make(chan Result, len(target))
	var wg sync.WaitGroup

	for i, t := range target {
		wg.Add(1)
		go binarySearch(arr, t, i, results, &wg)
	}

	wg.Wait()
	close(results)

	resultMap := make(map[int]int)
	for result := range results {
		resultMap[result.Index] = result.Value
	}

	for i, t := range target {
		result := resultMap[i]
		if result != -1 {
			fmt.Printf("Элемент %d найден на позиции %d\n", t, result)
		} else {
			fmt.Printf("Элемент %d не найден в массиве\n", t)
		}
	}
}
