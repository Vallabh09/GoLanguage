package main

import "fmt"

// func main(){
// 	result, err := divide(45,0)
// 	fmt.Println("division :", result)
// 	fmt.Println("error :", err)
// }

// func divide(a, b float64) (float64,error){
// 	if b == 0{
// 		return 0, fmt.Errorf("denominator cannot be zero")
// 	}else{
// 		return a/b, nil
// 	}
// }
//----------------------------------------------------------
// func main(){
// 	result, _ := divide(45,0)
// 	fmt.Println("division :", result)
// }

// func divide(a, b float64) (float64,error){
// 	if b == 0{
// 		return 0, fmt.Errorf("denominator cannot be zero")
// 	}else{
// 		return a/b, nil
// 	}
// }
//-------------------------------------------------------------
func main(){
	result, err := divide(45,0)
	fmt.Println("division :", result)
	fmt.Println("error :", err)
}

func divide(a, b float64) (float64, string){
	if b == 0{
		return 0, "denominator cannot be zero"
	}else{
		return a/b, "nil"
	}
}