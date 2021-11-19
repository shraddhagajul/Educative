package main

import "fmt"

func removeSpacesTabs(input string) string{
	//Base Case
	if len(input) == 0{
		return ""
	}
  //Recursive Case
	inputLength := len(input)

	if string(input[0]) ==  "\t" || string(input[0]) == " "{
		return removeSpacesTabs(input[1 : inputLength])
	}

	return string(input[0]) + removeSpacesTabs(input[1 : inputLength ])
}

func main(){
	result := removeSpacesTabs("\tHello World")
	fmt.Printf("Result : %s\n",result)
}