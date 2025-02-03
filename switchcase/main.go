package main

import "fmt"

func main(){
	n := 5
	switch n{
	case 1:
		fmt.Println("monday")
	case 2:
		fmt.Println("tuesday")
	case 3:
		fmt.Println("wednesday")
	default:
		fmt.Println("invalid input")
	}

	a := 0.99
	switch a{
	case 1, 3, 5, 7:
		fmt.Println("odd number")
	case 2, 4, 6, 8:
		fmt.Println("even number")
	case 0:
		fmt.Println("zero")
	default:
		fmt.Println("invalid input")
	}

	temperature := -10
	switch{
	case temperature < 0:
		fmt.Println("freezing")
	case temperature >= 0 && temperature < 10:
		fmt.Println("cold")
	case temperature >= 10 && temperature < 20:
		fmt.Println("cool")
	case temperature >= 20 && temperature < 30:
		fmt.Println("warm")
	default:
		fmt.Println("hot")
	}
}