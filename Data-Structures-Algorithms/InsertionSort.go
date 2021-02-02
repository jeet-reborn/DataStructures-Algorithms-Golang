package main

import "fmt"

func main() {
	arr := []int{5, 2, 3, 7, 9, 0}
	fmt.Println("Before Sorting : ", arr)
	sortedArray := InsertionSort(arr, len(arr))
	fmt.Println("After Sorting : ", sortedArray)
}

// InsertionSort : Insertion Sort function.
func InsertionSort(arr []int, n int) []int {

	for i := 1; i < n; i++ {
		j := i - 1
		x := arr[i] // x is the element that needs to be inserted at each pass.
		for j > -1 && arr[j] > x {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = x
	}
	return arr
}
