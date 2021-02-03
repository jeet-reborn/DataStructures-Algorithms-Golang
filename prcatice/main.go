package main

import "fmt"

func main() {
	var arr [100000]int

	for i := 0; i < len(arr); i++ {
		fmt.Scanln(&arr[i]++)
	}
}
