package main

import (
	"errors"
	"fmt"
)

func reverseStringIterative(input string)(string,error){
	var result string
	if len(input) <= 0{
		return "",errors.New("Empty String")
	}
	
	for _,character := range input{
		result =  string(character) + result
	}
	fmt.Printf("Result : %s\n",result)

	return result,nil
}

// func main(){
// 	reverseStringIterative("Hello World")

// }