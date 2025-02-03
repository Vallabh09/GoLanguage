package main

import "fmt"

func main() {

	var num1, num2 int 
	var operator string

	fmt.Print("enter first number :")
	fmt.Scanln(&num1)

	fmt.Print("enter second number :")
	fmt.Scanln(&num2)

	fmt.Print("enter operation :")
	fmt.Scanln(&operator)

	switch operator{

	case "+":
		fmt.Println(num1+num2)

	case "-":
		fmt.Println(num1-num2)

	case "*":
		fmt.Println(num1*num2)

	case "/":
		if num2 == 0{
			fmt.Println("divide by zero exception")
		}else{
			fmt.Println(num1/num2)
		}
	
	default:
		fmt.Println("invalid operator")
		
	}

}