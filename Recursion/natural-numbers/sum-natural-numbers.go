package main

import "fmt"

func sumNaturalNumbers(lowerRange,upperRange int){
	//Base Case
	if lowerRange > upperRange{
		return
	}

  //Recursion Case
	fmt.Printf(" LowerRange : %d\n",lowerRange)
	sumNaturalNumbers( lowerRange + 1, upperRange)

}

func main(){
	sumNaturalNumbers(1,5)
}
