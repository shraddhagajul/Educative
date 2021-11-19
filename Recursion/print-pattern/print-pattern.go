package main

import "fmt"

//Given a target number, subtract 5 from the targetNumber until we 
// reach a negative number or 0. Then we add 5 until we reach the 
// targetNumber again.

func printPattern(targetNumber int){

	//Base Case
	if (targetNumber <= 0 ){
		fmt.Printf("%d\n",targetNumber)
		return
	}

	fmt.Printf("%d\n",targetNumber)
	printPattern(targetNumber - 5)
	fmt.Printf("%d\n",targetNumber)

}

func main(){
	number := 15
	printPattern(number)
}