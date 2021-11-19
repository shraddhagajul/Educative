package main

import (
	"errors"
	"fmt"
)

func modulo(dividend,divisor int) (int,error){
	if divisor == 0{
		return 0,errors.New("Divisor should be greater than 0")
	}

	if divisor > dividend{
		return dividend,nil
	}

	return modulo(dividend - divisor,divisor)
}

func main(){
	result,err := modulo(14,7)
	if err != nil{
		fmt.Printf("Error : %s\n",err)	
	}

	fmt.Printf("Result : %d\n",result)

}