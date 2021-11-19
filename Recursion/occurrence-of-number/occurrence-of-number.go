package main

import "fmt"

func occurrenceOfNumber(input []int,key int) int{

	if len(input) == 0{
		return 0
	}

	if input[0] == key{
		return 1 + occurrenceOfNumber(input[1 : ],key)
	}
	
	return 0 + occurrenceOfNumber(input[1 : ],key)
	
}

func main() {
	input := []int{1,21,3,4,3,21,21,3,5,6,76,21}
	result := occurrenceOfNumber(input,2)
	fmt.Printf("Result : %d\n",result)
}