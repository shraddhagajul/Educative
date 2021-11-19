package main

import "fmt"

func removeAdjacentDuplicates(input string) string{
	//Base Case
	if len(input) <= 1{
		return input
	}

	//Recursive Case
	inputLength := len(input)
	
	if input[0] == input[1]{
		return removeAdjacentDuplicates(input[1 : inputLength])
	}

	return string(input[0]) + removeAdjacentDuplicates(input[1 : inputLength])

}

func main(){
	result := removeAdjacentDuplicates("")
	fmt.Printf("Result : %s\n",result)
}