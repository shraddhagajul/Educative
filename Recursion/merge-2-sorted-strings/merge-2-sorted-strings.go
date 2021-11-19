package main

import "fmt"

func mergeTowSortedString(input1, input2 string) string{
	input1Length := len(input1)
	input2Length := len(input2)

	if input1Length == 0 && input2Length != 0{
		return input2
	}
	if input1Length != 0 && input2Length == 0{
		return input1
	}

	if input1Length == 0 && input2Length == 0{
		return ""
	}

	// Recursive Case
	if input1[0] < input2[0]{
		return string(input1[0]) + mergeTowSortedString(input1[1 : input1Length] , input2)

	}
	return string(input2[0]) + mergeTowSortedString(input1 , input2[1 : input2Length])
}

func main(){
	result := mergeTowSortedString("acu" , "bst")
	fmt.Printf("Result : %s\n",result)
}