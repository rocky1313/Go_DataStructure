package main

import "fmt"

func main() {
	arr1 := [...]int{1, 4, 2, 4, 67, 8, 8, 9}
	fmt.Println(bubbleSort(arr1[:]))

}

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}

		}
	}
	return arr
}
