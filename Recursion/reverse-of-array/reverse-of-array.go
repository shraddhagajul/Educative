package main

import "fmt"

var result []int

func reverseOfArray(input []int) []int{

	if len(input) == 0{
		return result
	}

	result = append(result, input[len(input) - 1 ])

	return reverseOfArray(input[ : len(input) - 1])
}

func main() {
	input := []int{ 1, 2, 3, 4}
	result := reverseOfArray(input)

	fmt.Print(result)
}