package main

import ( "fmt" )

func decimalToBinary(input int) string{

  //Base Condition	
	if input == 0 {
		return ""
	}
	
  //Recursive 
	if input % 2 == 0{
		return decimalToBinary( input / 2) + "0"
	}

	return decimalToBinary( input / 2) + "1"

}

func modifyDecimalToBinary(input int) string{
	if input <= 1{
		return fmt.Sprint(input)
	}
	return modifyDecimalToBinary( int(input / 2 )) + modifyDecimalToBinary(input % 2) 
}

func main(){
	result := decimalToBinary(11)
	fmt.Printf("Result : %s\n",result)

	res := modifyDecimalToBinary(11)
	fmt.Printf("Result : %s\n",res)
}