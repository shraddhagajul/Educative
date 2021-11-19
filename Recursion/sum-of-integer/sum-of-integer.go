package main

import "fmt"

func sumOfIntegers(input int) int{
	//Base Case
	if input == 1{
		return input
	}

	return input + sumOfIntegers(input - 1)
}

func main(){

	result := sumOfIntegers(5)
	fmt.Printf("Result : %d\n",result)
	
}

