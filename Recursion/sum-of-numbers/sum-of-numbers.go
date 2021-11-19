package main

import (
	"fmt"
	"strconv"
)

func sumOfNumbers(input string) int {
	if input == ""{
		return 0
	}
	stringVal, err := strconv.ParseInt(string(input[0]),10,12)
	if err != nil{
		return -99999
	}
	fmt.Printf("input[0] : %d\n" ,stringVal )
	return int(stringVal) + sumOfNumbers(input[1 : ])
}

func main(){
	result := sumOfNumbers("345")
	fmt.Printf("Result : %d\n", result)
}