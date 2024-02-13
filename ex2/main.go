package main

import "fmt"


func EachCons(arr []int, n int) [][]int {
	// your code here
	var arr_main [][]int
	first_index := 0
	last_index := first_index + n

	for i:=0; last_index < len(arr)+1; i++ {
		arr_sub := arr[first_index: last_index]
		arr_main = append(arr_main, arr_sub)
		first_index += 1
		last_index += 1
	}
	return arr_main
}


func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(EachCons(arr, 2))
}
