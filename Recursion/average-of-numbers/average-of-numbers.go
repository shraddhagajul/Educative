package main

import "fmt"

var sum float64

func averageOfNumbers(input []int,currentIndex int) float64{
	//Base Case
	if len(input) == 0{
		return 0
	}

	if len(input) == currentIndex{
		return ( sum / float64(currentIndex) )
	}
	//Recursive Case
	sum += float64(input[currentIndex])
	return averageOfNumbers(input,currentIndex + 1)
}

func main(){
	result := averageOfNumbers([]int{1,2,3,4,5},0)
	fmt.Print(result)
}