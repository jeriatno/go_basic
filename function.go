package main

import "fmt"


var num1 = 10
var num2 = 5

/*
    single return
*/

// multiply
func multiply() int {
    return num1 * num2
}
// dividen
func dividen(num3 int, num4 int) int {
    return num3 / num4
}

/*
    multiple return
*/

func fullname(first_name string) (string, string) {
    return first_name, " Jeriatno";
}

func body() (weight string, height int) {
    weight = "70"
    height = 170
    return
}

func main() {
    // call to single return
    fmt.Printf("multiply = ", multiply(), "\n")
    fmt.Printf("dividen = ", dividen(30, 3), "\n\n")

    // call to multiple return
    first_name, last_name := fullname("iwan")
    fmt.Printf(first_name+" "+last_name+"\n")
    weight2, height2 := body()
    fmt.Printf(weight2, height2)
}
