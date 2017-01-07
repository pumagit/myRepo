package main

import "fmt"

func main() {
	numSlice := []int {5,4,3,2,1}


	// Slices with defined dimension
	numSlice3 := make([]int,5,10)
	copy(numSlice3,numSlice) // copy slice to slice
	fmt.Println(numSlice[0])

	// Append element to an slice
	numSlice3 = append(numSlice3,0,-2)
	fmt.Println(numSlice3[6])

}
