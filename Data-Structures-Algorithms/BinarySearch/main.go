// Implement Binary Search Algorithm.

package main

import "fmt"

func main() {
	itemToBeFound := 50
	theArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	length := len(theArray)

	found := recursiveBinarySearch(theArray, 0, length-1, itemToBeFound)
	if found == -1 {
		fmt.Printf("Element NOT found")
	} else {
		fmt.Printf("Element found at position %d", found)
	}

}

func recursiveBinarySearch(arr []int, l int, r int, item int) int {
	if r >= l {

		mid := (l + r) / 2

		if arr[mid] == item {
			return mid
		}
		if arr[mid] > item {
			return recursiveBinarySearch(arr, l, mid-1, item)
		}
		return recursiveBinarySearch(arr, mid+1, r, item)
	}
	return -1
}
