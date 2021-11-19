package main

import "fmt"

func isPalindrome(input string) bool {
	if len(input) <= 1 {
		return true
	}

	if string(input[0]) == input[len(input)-1:] {
		return true && isPalindrome(input[1:len(input)-1])
	}
	return false
}

func main() {
	result := isPalindrome("madam")
	fmt.Printf("Result : %v\n",result)
}