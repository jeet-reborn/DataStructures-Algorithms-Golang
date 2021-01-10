package main

import "fmt"

func main() {
	var arr = []int{5, 6, 3, 2, 8, 1, 0}
	fmt.Println("Before Sorting : ", arr)
	insertionSort(arr)
	fmt.Println("After Sorting : ", arr)
}

func insertionSort(arr []int) {

	for i := 0; i < len(arr); i++ {
		j := i
		key := arr[j]

		for j > 0 && key < arr[j-1] {

			arr[j] = arr[j-1]
			j = j - 1
		}
		arr[j] = key
	}
}
