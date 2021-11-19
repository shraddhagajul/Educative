package main

import "fmt"

func fibonacci(testVariable int) int{
	//Base Case
	if  testVariable <= 1{
		return testVariable
	}

	//Recursive Case
	return  ( fibonacci(testVariable - 1) + fibonacci( testVariable - 2 ) )
}

func main(){
	result := fibonacci(7)
	fmt.Printf("Result : %d\n",result)
}