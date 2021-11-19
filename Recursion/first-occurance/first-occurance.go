package main

import "fmt"

func firstOccurrence(arr []int, testVariable int, currentIndex int) int{
	if len(arr) == currentIndex {
		return -1
	}

	if arr[currentIndex] == testVariable{
		return currentIndex
	}

	currentIndex++
	
	return firstOccurrence(arr,testVariable,currentIndex )
}

func main(){
	arr := []int{10,20,30,40,50}
	testVariable := 80
	currentIndex := 0
	result := firstOccurrence(arr,testVariable,currentIndex)
	fmt.Printf("Result : %d\n",result)
}