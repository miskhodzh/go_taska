package main

import "fmt"


func MergeArrays(arr1, arr2 []int) []int {
	merged := append(arr1, arr2...)
	
	return merged
  }


func main() {
	arr1 := []int{1, 3, 5, 7}
	arr2 := []int{2, 4, 6, 8}
	fmt.Println(MergeArrays(arr1, arr2))
}
