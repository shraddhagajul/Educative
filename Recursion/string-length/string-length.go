package main

import "fmt"

func stringLength(input string ) int {
	//Base Case 
	if input == "" {
		return 0
	}
	return 1 + stringLength(input[ 1 : ])
}

func main(){
	result := stringLength("Educative")
	fmt.Printf("Result : %d\n",result)
}