package main

import "fmt"

func main() {
    // first step
    // for i:=0; i<5; i++ {
    //     fmt.Printf("Nilai i = ", i, "\n")
    // }

    // second step
    // i:=0;
    // for i<5 {
    //     fmt.Printf("Nilai i = ", i, "\n")
    //     i++
    // }

    // three step
    total := 0;
    for num:=0; num<=5; num++ {
        total = total + num
    }

    fmt.Printf("Total = ", total)
}
