package main

import "fmt"

func main(){
	a := 20
	if a > 10{
		fmt.Println("a is greater than 10")
	}else{
		fmt.Print("a is smaller than 10")
	}

	z := 200
	if z % 10 == 0 && z % 100 == 0{
		fmt.Println("z is multiple of 10 and 100")
	}else if z / 9 == 0{
		fmt.Println("z is multiple of 9")
	}else{
		fmt.Println("z is not divisible by 9 or 10")
	}

}