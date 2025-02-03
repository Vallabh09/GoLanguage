//slice is just a dynamic array, it changes its size during runtime
package main

import "fmt"

func main(){
	//general way of declaring and iniatlizing slice
	// var names = []string{}
	// names = append(names, "vallabh","joshi")
	// fmt.Println(len(names))
	// fmt.Println(cap(names))
	// fmt.Println(names)

	//we can create slice using make function
	numbers := make([]int,2,5)
	numbers = append(numbers,1,2,3,4,5,6,7,8,9,99,55,44,55,66,77,88)
	//numbers = append(numbers,7,8,9)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
	fmt.Println(numbers)

	//if we add new elements in same append then the capacity incremented by 2 each time u add
	//if we use new append statement the capacity is multiplyed by 2


}