package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	
	// //using fmt.scan()
	// var str string
	// fmt.Scan(&str)
	// fmt.Println("you typed :", str)

	// //multiple input
	// var str1, str2, str3 string
	// fmt.Scan(&str1, &str2, &str3)
	// fmt.Println("you typed :", str1, str2, str3)

	// //using fmt.Sprintf() to concat strings
	// var str4, str5, concat string
	// fmt.Scan(&str4, &str5)
	// concat = fmt.Sprintf("%s %s",str4, str5)
	// fmt.Println("you typed :", concat)

	// //using fmt.Scanf()
	// var state string
	// var passing int
	// var series string
	// var num int
	// fmt.Scanf("%2s%2d%2s%4d", &state, &passing, &series, &num)
	// carnum := fmt.Sprintf("%2s %2d %2s %4d",state,passing,series,num)
	// fmt.Println("your car no is :",carnum )
	// fmt.Println("state :", state)

	//using bufio reader read whole line unit \n, it creates buffer
	fmt.Print("enter a sentence :")

	reader := bufio.NewReader(os.Stdin)
	sentence, _ := reader.ReadString('\n')
	fmt.Println("your sentence :", sentence)
}