package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := []int{5, 2, 3, 1, 7, 8, 9, 5, 4, 3}
	fmt.Println("Before Sorting :")
	printArray(arr)

	QuickSort(arr, 0, len(arr)-1)
	fmt.Println("After Sorting :")
	printArray(arr)
}

//QuickSort : Pivot
func QuickSort(arr []int, l, h int) {
	if l < h {
		var pivotElement = Partition(arr, l, h)
		QuickSort(arr, l, pivotElement)
		QuickSort(arr, pivotElement+1, h)
	}
}

//Partition : Return the pivot.
func Partition(arr []int, l, h int) int {
	var pivotElement = arr[l]
	var i = l
	var j = h

	for i < j {
		for arr[i] <= pivotElement && i < h {
			i++
		}

		for arr[j] > pivotElement && j > l {
			j--
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[l] = arr[j]
	arr[j] = pivotElement

	return j
}

//printArray : Print the array.
func printArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(strconv.Itoa(arr[i]) + " ")
	}
	fmt.Println("")
}
