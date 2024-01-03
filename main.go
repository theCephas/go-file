package main

import "fmt"

func main(){

	var ages [3]int = [3]int{1, 2, 3}
	
	fmt.Println(ages, len(ages))

	//  Slices
	var scores = []int{1, 2, 3, 4, 5}
	scores = append(scores, 8)

	fmt.Println(scores, len(scores))

	// Slice ranges
	words := [6]string{"bells", "shoes", "fan", "cloth", "legs"}
	rangeOne := words[3:5]
	// Result is [cloth legs]
	rangeOne = words[2:5]
	// Result is [fan legs]
	
	fmt.Println(rangeOne)
}
