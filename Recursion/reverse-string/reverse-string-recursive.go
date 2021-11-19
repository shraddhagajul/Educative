package main

import "fmt"

func reverseStringRecursive(input string){
	//Base Case
	if len(input) <= 0{
		return
	}
	// Recursive Case
	inputLength := len(input)
  fmt.Printf("%v",input[inputLength - 1 : ])
	reverseStringRecursive(input[ :inputLength - 1])
	

}

func main(){
	reverseStringRecursive("Hello World")
	fmt.Println("")
	reverseStringRecursive("Shraddha")
	fmt.Println("")
}

