package main

import "fmt"

func main() {
	arr := []int{5, 3, 6, 8, 9, 1, -2, -1, -50, 25, 29, 11111, -5990}
	fmt.Println("Before Sorting : ", arr)
	sortedArray := SelectionSort(arr, len(arr))
	fmt.Println("After Sorting : ", sortedArray)
}

//SelectionSort : Function to sort.
func SelectionSort(arr []int, n int) []int {

	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}
