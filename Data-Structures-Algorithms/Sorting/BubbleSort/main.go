package main

import "fmt"

func main() {
	arr := []int{5, 6, 9, 2, 4, 1}
	fmt.Println("Before Sorting : ", arr)
	sortedArray := BubbleSort(arr, len(arr))
	fmt.Println("After Sorting : ", sortedArray)
}

// BubbleSort : Function to sort.
func BubbleSort(arr []int, n int) []int {

	for i := 0; i < n-1; i++ {
		swapped := 0
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j] //Swap the elements
				swapped = 1
			}
		}
		if swapped == 0 {
			break
		}
	}
	return arr
}
