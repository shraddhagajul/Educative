package main

import "fmt"

func balanceParenthesis(input []byte,startIndex,currentIndex int) bool{
	if len(input) == startIndex{
		if currentIndex == 0{
			return true 
		}
		return false
	}

	if currentIndex < 0{
		return false
	}

	if input[startIndex] == '('{
		return balanceParenthesis(input,startIndex+1,currentIndex+1)
	}

	if input[startIndex] == ')'{
		return balanceParenthesis(input,startIndex+1,currentIndex-1)
	}

	return false
}

func main(){
	result := balanceParenthesis([]byte{'(','(',')',')'},0,0)
	fmt.Printf("Result : %v\n",result)
}