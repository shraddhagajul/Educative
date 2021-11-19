package main

import "fmt"

func rowValue(y,x int) int{
	if y < 0 || x < 0 {
		return 0
	}

	if y >= 0 && x == 0 {
		return 1       
	}
	return rowValue(y - 1,x - 1) + rowValue(y - 1,x)
}

func pascalsTriangle(input int) []int{
	result := make([]int,input+1)

	 for i := 0; i<= input ; i++{
		result[i] = rowValue(input, i)
	 }
	return result
}

func main(){
	printResult := pascalsTriangle(5)
	// printResult := rowValue(1, 1)
	fmt.Print(printResult)
}

