package main

import "fmt"

func replaceNegativeNumber(input []int,currentIndex int) []int{
	if currentIndex == len(input){
		return input
	}
	if input[currentIndex] < 0{
		input[currentIndex] = 0
 	}
	return replaceNegativeNumber(input,currentIndex + 1)
}

func main(){
	result := replaceNegativeNumber([]int{2, -3, 4, -1, -7, 8},0)
	fmt.Print(result)
}