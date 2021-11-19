package main

import "fmt"


func powerOfNumber(base,exponent int) int{

	if exponent == 0{
		return 1
	}
	return base * powerOfNumber( base,exponent - 1 )
}

func main(){
	result := powerOfNumber(2,5)
	fmt.Printf("Result : %d\n",result)
}