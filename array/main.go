package main

import "fmt"

func main() {
	var numbers [5]float64
	numbers[0] = 10.00
	numbers[2] = 20.35
	numbers[4] = 40.77

	fmt.Println("numbers in array:", numbers)

	fmt.Println("number at index 0:", numbers[0])
	fmt.Println("number at index 1:", numbers[1])
	fmt.Println("number at index 2:", numbers[2])
	fmt.Println("number at index 4:", numbers[4])

	var names = [5] string {"merc","bmw","aston","rr","bentley"}
	fmt.Println("names in array :", names)
	fmt.Println("length of numbers array:", len(numbers))

	fmt.Println("length of names array:", len(names))

	//default values for empty element
	//int:0
	//float:0
	//string:""
	//bool:false
}
