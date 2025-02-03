package main

import "fmt"

func main() {
    // hello world
    fmt.Println("hello world")

    // variables
    var num int = 10
    var decimal float64 = 20.25
    var str string = "vallabh"
    var result bool = true

    fmt.Println(num)
    fmt.Println(decimal)
    fmt.Println(str)
    fmt.Println(result)

    var a = 20
    fmt.Println(a)

    a = 40
    fmt.Println(a)

    //print statements
    fmt.Print("this is print function\n")
    fmt.Println("this is println function, it prints the string and move cursor on next line.")
    fmt.Printf("this is printf function, it prints variables with their format.\n")

    fmt.Printf("this is integer a %d.\n", a)
    fmt.Printf("thus is float %.2f.\n", decimal)
    fmt.Printf("this is string %s.\n", str)
    fmt.Printf("this returns a datatype of result : %T.\n", result)

}

