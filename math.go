package main

import "fmt"


func main() {
    var num1 int
    var num2 int

    num1 = 10;
    num2 = 20;
    // increase
    increase := 10 + 5
    // decrease
    decrease := 10 - 5
    // increments
    num1++
    // decrements
    num2--
    // multiply
    multiply := increase * decrease
    // dividen
    dividen := increase / decrease
    // mods
    mods := increase % decrease

    fmt.Printf("increase = ", increase,
                "\n decrease = ", decrease,
                "\n increments = ", num1,
                "\n decrements = ", num2,
                "\n multiply = ", multiply,
                "\n dividen = ", dividen,
                "\n mods = ", mods)
}
