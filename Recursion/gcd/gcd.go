package main

func gcd(input1,input2 int) int{
	if input1 == input2{
		return input1
	}

	if input1 > input2{
		return gcd(input1-input2,input2)
	}

	return gcd(input1, input2-input1)
}

