package main

import (
	"fmt"
)

func countVowelsRecursive(input string,counter int)(int){
	if len(input) == 0{
		return counter
	} 

	if (isVowel(byte(input[0]))){
		counter++
	}
	
	 return countVowelsRecursive(input[ 1 : ],counter)

}	

func isVowel(element byte) bool{
	vowels := []rune{'a','e','i','o','u'}

	for _, vowel := range vowels{
		if rune(element) == vowel{
			return true
		}
	}
		return false	
}

func main(){
  count := countVowelsRecursive("Educative", 0)
	fmt.Printf("Vowels count : %d\n",count)

}