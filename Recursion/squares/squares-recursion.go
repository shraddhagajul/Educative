package main

import "fmt"

func squares(testVariable int) int{

		if testVariable == 0 {
			return 0
		}
		
		return squares(testVariable - 1) + ( 2 * testVariable ) - 1 

}

func main() {
	result := squares(5)
	fmt.Printf("Result : %d\n",result)

}