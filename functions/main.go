package main

import "fmt"

func main(){
	result := add(10,20)
	fmt.Println("addition :", result)

	simplefunc()
}

func add(a, b int) (c int){
	c = a+b
	return

}

func simplefunc(){
	fmt.Println("this is simple function with no arguments.")
}