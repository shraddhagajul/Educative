package main

import "fmt"

func Factorial(targetNumber int) int{

	//Base condition
	if targetNumber <= 1{
		return 1
	}
	// Recursion Condition
	return ( targetNumber * Factorial(targetNumber - 1) )
}

func main(){
	var result = Factorial(5)
	fmt.Printf("Result : %d\n",result)
}