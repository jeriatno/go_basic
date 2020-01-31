package main

import "fmt"

func main() {
    num := 50;
    col := "silver"

    // first step
    if num > 10 {
        fmt.Printf("Num > 10 \n")
    } else {
        fmt.Printf("Num < 10 \n")
    }

    // second step
    if col == "silver" {
        fmt.Printf("Color = Silver \n")
    } else {
        fmt.Printf("Color != Silver \n")
    }

    // three step
    if true {
        fmt.Printf("You are smart \n")
    } else {
        fmt.Printf("You are not smart \n")
    }

    // four step
    money := 5000
    loan  := 5000

    if money > loan {
        fmt.Printf("You maney is more \n")
    } else if money == loan {
        fmt.Printf("You money is enough \n")
    } else {
        fmt.Printf("You money is less \n")
    }
}
